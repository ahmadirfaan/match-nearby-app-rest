package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/sirupsen/logrus"
)

type UserManageUsecase interface {
	UpdateProfile(request web.UpdateProfileRequest) error
}

type userManageUsecase struct {
	userRepository repositories.UsersRepository
}

func NewUserManageUsecase(ur repositories.UsersRepository) UserManageUsecase {
	return &userManageUsecase{
		userRepository: ur,
	}
}

func (um *userManageUsecase) UpdateProfile(request web.UpdateProfileRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		logrus.Info("error validators: " + err.Error())
		return utils.ErrorValidator
	}

	return nil
}
