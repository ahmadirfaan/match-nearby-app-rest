package utils

import (
	"github.com/sirupsen/logrus"
	"time"

	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/golang-jwt/jwt"
)

type ConfigToken struct {
	JWTSecret string
	TokenTTL  int
}

func GenerateToken(user database.Users) (*string, *uint64, error) {

	configApp := config.Init()
	jwtSecretKey := configApp.JWTSecret
	configExpired := configApp.TokenTTL

	expiredTime := time.Unix(time.Now().UTC().Add(time.Duration(configExpired)*time.Second).Unix(), 0)

	// Generates Access Token Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["userId"] = user.ID
	claims["exp"] = expiredTime
	logrus.Infof("jwt secret key from utils: %v", jwtSecretKey)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return nil, nil, err
	}

	timePointer := uint64(expiredTime.UnixMilli() / 1000)
	return &t, &timePointer, err
}
