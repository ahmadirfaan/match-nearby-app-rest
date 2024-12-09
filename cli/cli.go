package cli

import (
	"context"
	"fmt"
	"github.com/ahmadirfaan/match-nearby-app-rest/config/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/ahmadirfaan/match-nearby-app-rest/middleware"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/routes"
	"github.com/ahmadirfaan/match-nearby-app-rest/usecase"
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

	//setup the connection
	db := storage.InitDb()

	//create repository
	userRepository := repositories.NewUserRepository(db)
	profileRepository := repositories.NewProfileRepository(db)

	usecase.NewUserAuthenticationUsecase(userRepository, profileRepository)

	//create each use case
	userAuthenticationUsecase := usecase.NewUserAuthenticationUsecase(userRepository, profileRepository)
	userManageUsecase := usecase.NewUserManageUsecase(userRepository, profileRepository)

	//create routes
	authRoutes := routes.NewAuthRoutes(userAuthenticationUsecase)
	userRoutes := routes.NewUserRoutes(userManageUsecase)

	ginApp := gin.Default()
	configMiddleware(ginApp)

	prefixApiURL := "/api/v1"
	authMiddleware := middleware.AuthMiddlewareJWT()

	//create group auth
	authGroup := ginApp.Group(prefixApiURL + "/auth")
	{
		authGroup.POST("/signup", authRoutes.SignUp)
		authGroup.POST("/login", authRoutes.SignIn)
	}

	userGroup := ginApp.Group(prefixApiURL + "/users")
	{
		userGroup.PUT("", authMiddleware, userRoutes.UpdateProfile)
	}

	swipeGroup := ginApp.Group(prefixApiURL + "/swipes")
	swipeGroup.Use(authMiddleware)
	{
		swipeGroup.POST("", authMiddleware, userRoutes.UpdateProfile)
	}

	subscriptionsGroup := ginApp.Group(prefixApiURL + "/subscriptions")
	subscriptionsGroup.Use(authMiddleware)
	{
		subscriptionsGroup.POST("", authMiddleware, userRoutes.UpdateProfile)
	}

	StartServerWithGracefulShutdown(ginApp, app.Config)
}

func configMiddleware(ginApp *gin.Engine) {
	ginApp.Use(middleware.ErrorHandler())
	ginApp.NoRoute(middleware.NoRouteHandler)
	ginApp.NoMethod(middleware.NoRouteHandler)
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

	log.Printf("Server running on port :%s", serverconfig.AppPort)

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // Listen for SIGINT, SIGTERM
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
