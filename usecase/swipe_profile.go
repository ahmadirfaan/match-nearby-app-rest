package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/config/storage"
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

	user := u.userRepository.GetByUserId(userID)

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

		if totalSwipe <= 0 {
			redisClient.Set(ctxBackgroundRedis, keySwipe, "1", 24*time.Hour)
		}

		redisClient.Incr(ctxBackgroundRedis, keySwipe)

	}

	return nil

}
