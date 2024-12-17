package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
)

type SwipeUsecase interface {
}

type swipeUsecase struct {
	userRepository         repositories.UsersRepository
	profileRepository      repositories.ProfilesRepository
	subscriptionRepository repositories.SubscriptionsRepository
}

func NewSwipeUseCase(ur repositories.UsersRepository, pr repositories.ProfilesRepository, sr repositories.SubscriptionsRepository) SwipeUsecase {
	return &swipeUsecase{
		userRepository:         ur,
		profileRepository:      pr,
		subscriptionRepository: sr,
	}
}
