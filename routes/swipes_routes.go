package routes

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/usecase"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/gin-gonic/gin"
)

type SwipeRoutes interface {
	SwipeAction(context *gin.Context)
	GetNearbyProfile(context *gin.Context)
}

type swipeRoutes struct {
	SwipeUsecase usecase.SwipeUsecase
}

func NewSwipeRoutes(uc usecase.SwipeUsecase) SwipeRoutes {
	return swipeRoutes{
		SwipeUsecase: uc,
	}
}

func (sr swipeRoutes) SwipeAction(c *gin.Context) {
	var request web.SwipeRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.Error(utils.ErrorBadRequest)
		return
	}

	if err := sr.SwipeUsecase.SwipeProfiles(c.GetString("userID"), request); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Swipe Action",
	})
}

func (sr swipeRoutes) GetNearbyProfile(c *gin.Context) {

	data, remainingQuota, err := sr.SwipeUsecase.GetProfiles(c.GetString("userID"))

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, &web.GetProfileResponse{
		Data:           data,
		RemainingQuota: *remainingQuota,
	})
}
