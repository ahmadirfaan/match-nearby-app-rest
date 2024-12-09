package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
)

type UserManageUsecase interface {
	UpdateProfile(userID string, request web.UpdateProfileRequest) error
}

type userManageUsecase struct {
	userRepository    repositories.UsersRepository
	profileRepository repositories.ProfilesRepository
}

func NewUserManageUsecase(ur repositories.UsersRepository, pr repositories.ProfilesRepository) UserManageUsecase {
	return &userManageUsecase{
		userRepository:    ur,
		profileRepository: pr,
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

	return um.profileRepository.SaveProfile(profile)
}
