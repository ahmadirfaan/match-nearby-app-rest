package utils

import (
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/golang-jwt/jwt"
)

type ConfigToken struct {
	JWTSecret string
	TokenTTL  int
}

func GenerateToken(user database.Users) (string, *int64, error) {

	configApp := config.Init()
	jwtSecretKey := configApp.JWTSecret
	configExpired := configApp.TokenTTL

	expiredTime := time.Now().UTC().Add(time.Duration(configExpired) * time.Second).Unix()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expiredTime,
		"iat":     time.Now().Unix(),
		"sub":     "auth",
	}

	log.Info("jwtSecretKey: " + jwtSecretKey)

	// Create a new token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	t, err := token.SignedString([]byte(jwtSecretKey))
	log.Info("token generated: " + t)
	if err != nil {
		return "", nil, err
	}
	return t, &expiredTime, err
}
