package routes

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/usecase"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/gin-gonic/gin"
)

type UserRoutes interface {
	UpdateProfile(c *gin.Context)
}

type userRoutes struct {
	UserUsecase usecase.UserManageUsecase
}

func NewUserRoutes(uc usecase.UserManageUsecase) UserRoutes {
	return userRoutes{
		UserUsecase: uc,
	}
}

func (ar userRoutes) UpdateProfile(c *gin.Context) {
	var request web.UpdateProfileRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(utils.ErrorBadRequest)
		return
	}

	if err := ar.UserUsecase.UpdateProfile(request); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Update Profile",
	})
}
