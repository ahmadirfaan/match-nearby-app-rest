package main

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/gin-gonic/gin"
)

func main() {

	config, _ := config.LoadConfig()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	router.Run(":" + config.APIPort)
}
