package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/config/storage"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var redisClient = storage.InitRedis()

type SwipeUsecase interface {
	SwipeProfiles(userID string, request web.SwipeRequest) error
}

type swipeUsecase struct {
	userRepository  repositories.UsersRepository
	swipeRepository repositories.SwipeRepository
}

func NewSwipeUseCase(ur repositories.UsersRepository, sr repositories.SwipeRepository) SwipeUsecase {
	return &swipeUsecase{
		userRepository:  ur,
		swipeRepository: sr,
	}
}

func (u *swipeUsecase) SwipeProfiles(userID string, request web.SwipeRequest) error {

	if err := utils.NewValidator().Struct(&request); err != nil {
		return utils.ErrorValidator
	}

	user := u.userRepository.GetByUserId(userID)

	//can't self swipe
	if user.Profile.ID == request.UserID {
		return utils.ErrorBadRequest
	}

	keySwipe := fmt.Sprintf("SWIPE_%s", userID)

	if !user.IsPremium {
		ctxBackgroundRedis := context.Background()
		totalSwipeRedis := redisClient.Get(ctxBackgroundRedis, keySwipe)
		if err := totalSwipeRedis.Err(); !errors.Is(err, redis.Nil) && err != nil {
			return err
		}

		totalSwipeString, _ := totalSwipeRedis.Result()
		totalSwipe, _ := strconv.Atoi(totalSwipeString)

		if totalSwipe >= app.Init().Config.MaximalTotalSwipe {
			return utils.ErrorBadRequest
		}

		err := saveSwipe(userID, request, u)
		if err != nil {
			return err
		}

		if totalSwipe <= 0 {
			redisClient.Set(ctxBackgroundRedis, keySwipe, "1", 24*time.Hour)
		} else {
			redisClient.Incr(ctxBackgroundRedis, keySwipe)
		}

		return nil

	}

	return saveSwipe(userID, request, u)

}

func saveSwipe(userID string, request web.SwipeRequest, u *swipeUsecase) error {
	swipeDirection := "Pass"
	if request.Action {
		swipeDirection = "Like"
	}

	swipe := &database.Swipes{
		SwiperID:  userID,
		Direction: swipeDirection,
		SwipedID:  request.UserID,
		SwipedAt:  time.Now(),
	}

	if err := u.swipeRepository.SaveSwipe(swipe); err != nil {
		return err
	}
	return nil
}
