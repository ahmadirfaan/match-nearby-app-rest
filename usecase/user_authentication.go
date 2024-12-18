package usecase

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/repositories"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
)

type UserAuthenticationUseCase interface {
	Register(request web.SignUpRequest) error
	SignIn(request web.SignInRequest) (*web.SignInResponse, error)
	CheckUserExist(userID string) bool
}

type userAuthentication struct {
	userRepository    repositories.UsersRepository
	profileRepository repositories.ProfilesRepository
}

func NewUserAuthenticationUsecase(ur repositories.UsersRepository, pr repositories.ProfilesRepository) UserAuthenticationUseCase {
	return &userAuthentication{
		userRepository:    ur,
		profileRepository: pr,
	}
}

func (userAuth *userAuthentication) Register(request web.SignUpRequest) error {
	if err := utils.NewValidator().Struct(&request); err != nil {
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

	return userAuth.profileRepository.SaveProfile(profile)
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
		AccessToken: token,
		ExpiredAt:   *expire,
	}, nil
}

func (userAuth *userAuthentication) CheckUserExist(userID string) bool {
	if userID == "" {
		return false
	}
	user := userAuth.userRepository.GetByUserId(userID)
	return user != nil
}
