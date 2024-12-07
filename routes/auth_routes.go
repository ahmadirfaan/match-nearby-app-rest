package routes

import "github.com/gin-gonic/gin"

func SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Sign Up!",
	})
}

func SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Sign In!",
	})
}
