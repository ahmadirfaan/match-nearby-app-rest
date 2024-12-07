package cli

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/gin-gonic/gin"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (cli *Cli) Run(app *app.Application) {
	ginApp := gin.Default()

	ginApp.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	StartServerWithGracefulShutdown(ginApp, app.Config)
}

func StartServerWithGracefulShutdown(ginEngine *gin.Engine, serverconfig *config.Config) {
	// Configure the server
	srv := &http.Server{
		Addr:    fmt.Sprintf(`:%s`, serverconfig.AppPort),
		Handler: ginEngine,
	}

	// Run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println(fmt.Sprintf("Server running on port :%s", serverconfig.AppPort))

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill) // Listen for SIGINT, SIGTERM
	<-quit

	log.Println("Shutting down server...")

	// Create a context with timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(serverconfig.AppTimeout)*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Oops... Server is not shutting down! Reason: %v", err)
	}

	log.Println("Server exiting")

}
