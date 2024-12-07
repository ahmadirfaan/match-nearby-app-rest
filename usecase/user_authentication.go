package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
)

type userAuthentication struct {
	userRepository repositories.UsersRepository
}

type UserAuthenticationUseCase interface {
	Register()
}

func NewUserAuthenticationUsecase(ur repositories.UsersRepository) UserAuthenticationUseCase {
	return &userAuthentication{
		userRepository: ur,
	}
}

func (userAuth *userAuthentication) Register() {

}
