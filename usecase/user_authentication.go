package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/sirupsen/logrus"
)

type userAuthentication struct {
	userRepository repositories.UsersRepository
}

type UserAuthenticationUseCase interface {
	Register(request web.SignUpRequest) error
}

func NewUserAuthenticationUsecase(ur repositories.UsersRepository) UserAuthenticationUseCase {
	return &userAuthentication{
		userRepository: ur,
	}
}

func (userAuth *userAuthentication) Register(request web.SignUpRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		logrus.Info("error validators: " + err.Error())
		return utils.ErrorValidator
	}

	return nil
}
