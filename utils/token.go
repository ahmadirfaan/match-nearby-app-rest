package utils

import (
	"errors"
	"sync"
	"time"

	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/golang-jwt/jwt"
)

type ConfigToken struct {
	JWTSecret string
	TokenTTL  int
}

var cachedConfig *ConfigToken
var configOnce sync.Once

// This function to generates Access Token
func GenerateToken(user database.Users) (*string, *uint64, error) {

	var err error
	configOnce.Do(func() {
		// Only load the config once
		cfg := config.Init()
		cachedConfig = &ConfigToken{
			JWTSecret: cfg.JWTSecret,
			TokenTTL:  cfg.TokenTTL,
		}
	})

	if cachedConfig == nil {
		return nil, nil, errors.New("configuration not initialized")
	}
	jwtSecretKey := cachedConfig.JWTSecret
	configExpired := cachedConfig.TokenTTL

	time := time.Unix(time.Now().UTC().Add(time.Duration(configExpired)*time.Second).Unix(), 0)

	profile := user.Profile
	// Generates Access Token Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["userId"] = user.ID
	claims["profile"] = web.ProfileResponse{
		Bio:      profile.Bio,
		Gender:   profile.Gender,
		PhotoURL: profile.PhotoURL,
		Name:     profile.Name,
	}
	claims["exp"] = time
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return nil, nil, err
	}

	timePointer := uint64(time.UnixMilli() / 1000)
	return &t, &timePointer, err
}

// func ExtractToken(c *fiber.Ctx) (string, error) {
// 	appFiber := app.Init()
// 	tokenString := c.Get("Authorization")
// 	tokenString = strings.ReplaceAll(tokenString, " ", "")
// 	tokenString = strings.ReplaceAll(tokenString, "Bearer", "")
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(appFiber.Config.JWTSecret), nil
// 	})
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := claims["userId"].(float64)
// 	return strconv.Itoa(int(userId)), err
// }
