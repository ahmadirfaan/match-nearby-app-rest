package routes

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/usecase"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/gin-gonic/gin"
)

type AuthRoutes interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type authRoutes struct {
	AuthUseCase usecase.UserAuthenticationUseCase
}

func NewAuthRoutes(auc usecase.UserAuthenticationUseCase) AuthRoutes {
	return authRoutes{
		AuthUseCase: auc,
	}
}

func (ar authRoutes) SignUp(c *gin.Context) {
	var request web.SignUpRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(utils.ErrorBadRequest)
		return
	}

	if err := ar.AuthUseCase.Register(request); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Register",
	})
}

func (ar authRoutes) SignIn(c *gin.Context) {
	var request web.SignInRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(utils.ErrorBadRequest)
		return
	}

	response, err := ar.AuthUseCase.SignIn(request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, &response)
}
