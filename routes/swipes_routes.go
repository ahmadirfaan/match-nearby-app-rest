package routes

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/usecase"
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

func (sr swipeRoutes) SwipeAction(context *gin.Context) {

}

func (sr swipeRoutes) GetNearbyProfile(context *gin.Context) {

}
