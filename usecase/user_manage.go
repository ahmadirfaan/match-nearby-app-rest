package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"time"
)

type UserManageUsecase interface {
	UpdateProfile(userID string, request web.UpdateProfileRequest) error
	UpdatePremium(userID string) error
}

type userManageUsecase struct {
	userRepository         repositories.UsersRepository
	profileRepository      repositories.ProfilesRepository
	subscriptionRepository repositories.SubscriptionsRepository
}

func NewUserManageUsecase(ur repositories.UsersRepository, pr repositories.ProfilesRepository, sr repositories.SubscriptionsRepository) UserManageUsecase {
	return &userManageUsecase{
		userRepository:         ur,
		profileRepository:      pr,
		subscriptionRepository: sr,
	}
}

func (um *userManageUsecase) UpdateProfile(userID string, request web.UpdateProfileRequest) error {
	if err := utils.NewValidator().Struct(&request); err != nil {
		return utils.ErrorValidator
	}

	user := um.userRepository.GetByUserId(userID)
	if user == nil {
		return utils.ErrorForbidden
	}

	profile := &user.Profile
	if request.Name != "" {
		profile.Name = request.Name
	}

	if request.Bio != "" {
		profile.Bio = request.Bio
	}

	if request.PhotoURL != "" {
		profile.PhotoURL = request.PhotoURL
	}

	if request.Gender != "" {
		profile.Gender = request.Gender
	}

	return um.userRepository.SaveUser(user)
}

func (um *userManageUsecase) UpdatePremium(userID string) error {
	user := um.userRepository.GetByUserId(userID)
	if user == nil {
		return utils.ErrorForbidden
	}

	if user.IsPremium {
		return utils.ErrorBadRequest
	}

	user.IsPremium = true
	now := time.Now().UTC()

	premiumExpiry := now.AddDate(1, 0, 0)
	user.PremiumExpiry = &premiumExpiry
	user.IsPremium = true

	if err := um.userRepository.SaveUser(user); err != nil {
		return err
	}

	subscription := database.Subscriptions{
		UserID:       user.ID,
		PurchaseName: "Premium Swipes",
		PurchaseDate: now,
	}

	return um.subscriptionRepository.SaveSubscription(&subscription)
}
