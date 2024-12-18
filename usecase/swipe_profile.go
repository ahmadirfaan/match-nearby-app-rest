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
	"math"
	"strconv"
	"time"
)

var redisClient = storage.InitRedis()

var keySwipe = func(userID string) string {
	return fmt.Sprintf("SWIPE_%s", userID)
}

var maximalTotalSwipe = app.Init().Config.MaximalTotalSwipe

type SwipeUsecase interface {
	SwipeProfiles(userID string, request web.SwipeRequest) error
	GetProfiles(s string) ([]web.ProfileModelResponse, *uint16, error)
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

func (u *swipeUsecase) GetProfiles(userID string) ([]web.ProfileModelResponse, *uint16, error) {

	user := u.userRepository.GetByUserId(userID)

	var remainingQuota uint16
	if user.IsPremium {
		remainingQuota = math.MaxUint16
	} else {
		totalSwipeQuotaRedis := redisClient.Get(context.Background(), keySwipe(userID))
		if err := totalSwipeQuotaRedis.Err(); !errors.Is(err, redis.Nil) && err != nil {
			return nil, nil, err
		}

		totalSwipeString, _ := totalSwipeQuotaRedis.Result()
		totalSwipe, _ := strconv.Atoi(totalSwipeString)

		if totalSwipe >= maximalTotalSwipe {
			return nil, nil, utils.ErrorBadRequest
		}
		remainingQuota = uint16(maximalTotalSwipe - totalSwipe)
	}

	profiles := u.swipeRepository.GetSwipeNearby(userID)

	return convertProfilesToProfileModels(profiles), &remainingQuota, nil

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

	if !user.IsPremium {
		ctxBackgroundRedis := context.Background()
		totalSwipeRedis := redisClient.Get(ctxBackgroundRedis, keySwipe(userID))
		if err := totalSwipeRedis.Err(); !errors.Is(err, redis.Nil) && err != nil {
			return err
		}

		totalSwipeString, _ := totalSwipeRedis.Result()
		totalSwipe, _ := strconv.Atoi(totalSwipeString)

		if totalSwipe >= maximalTotalSwipe {
			return utils.ErrorBadRequest
		}

		err := saveSwipe(userID, request, u)
		if err != nil {
			return err
		}

		if totalSwipe <= 0 {
			redisClient.Set(ctxBackgroundRedis, keySwipe(userID), "1", 24*time.Hour)
		} else {
			redisClient.Incr(ctxBackgroundRedis, keySwipe(userID))
		}

		return nil

	}

	return saveSwipe(userID, request, u)

}

func saveSwipe(userID string, request web.SwipeRequest, u *swipeUsecase) error {

	//check swipe action between one day
	swipeData := u.swipeRepository.GetSwipeStatus(userID, request.UserID)
	if len(swipeData) > 0 {
		return utils.ErrorBadRequest
	}

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

func convertProfilesToProfileModels(profiles []database.Profiles) []web.ProfileModelResponse {
	var profileModels []web.ProfileModelResponse

	for _, p := range profiles {
		profileModels = append(profileModels, web.ProfileModelResponse{
			UserId: p.UserID,
			Name:   p.Name,
			Gender: p.Gender,
			Photo:  p.PhotoURL,
			Bio:    p.Bio,
		})
	}

	return profileModels
}
