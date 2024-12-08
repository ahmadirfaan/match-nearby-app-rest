package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/sirupsen/logrus"
)

type userAuthentication struct {
	userRepository    repositories.UsersRepository
	profileRepository repositories.ProfilesRepository
}

type UserAuthenticationUseCase interface {
	Register(request web.SignUpRequest) error
	SignIn(request web.SignInRequest) (*web.SignInResponse, error)
}

func NewUserAuthenticationUsecase(ur repositories.UsersRepository, pr repositories.ProfilesRepository) UserAuthenticationUseCase {
	return &userAuthentication{
		userRepository:    ur,
		profileRepository: pr,
	}
}

func (userAuth *userAuthentication) Register(request web.SignUpRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		logrus.Info("error validators: " + err.Error())
		return utils.ErrorValidator
	}

	//create user
	user := &database.Users{
		Email:    request.Email,
		Password: utils.HashPassword(request.Password),
		Username: request.Username,
	}
	if err := userAuth.userRepository.SaveUser(user); err != nil {
		return err
	}

	//create profile
	profile := &database.Profiles{
		UserID: user.ID,
		Gender: request.Gender,
		Name:   request.Name,
	}
	if err := userAuth.profileRepository.SaveProfile(profile); err != nil {
		return err
	}

	return nil
}

func (userAuth *userAuthentication) SignIn(request web.SignInRequest) (*web.SignInResponse, error) {
	err := utils.NewValidator().Struct(&request)
	isEmailEmpty := request.Email == ""
	isUsernameEmpty := request.Username == ""
	isEmailAndUsernameEmpty := isEmailEmpty && isUsernameEmpty

	if err != nil || isEmailAndUsernameEmpty {
		return nil, utils.ErrorValidator
	}

	var user *database.Users
	if isEmailEmpty {
		user = userAuth.userRepository.GetByUsername(request.Username)
	} else {
		user = userAuth.userRepository.GetByEmail(request.Email)
	}

	if user == nil {
		return nil, utils.ErrorNotFound
	}
	isPasswordEquals := utils.CheckPasswordHash(request.Password, user.Password)

	if !isPasswordEquals {
		return nil, utils.ErrorAuth
	}

	token, expire, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &web.SignInResponse{
		TokenType:   "Bearer",
		AccessToken: *token,
		ExpiredAt:   *expire,
	}, nil
}
