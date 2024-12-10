package middleware

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/usecase"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/ahmadirfaan/match-nearby-app-rest/config"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/gin-gonic/gin"
)

var user_id_key = "userID"
var user_id_key_claim = "user_id"

func AuthMiddlewareJWT(userUseCase usecase.UserAuthenticationUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logrus.Errorf("Header blank")
			c.Error(utils.ErrorAuth)
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logrus.Errorf("error Bearer not valid")
			c.Error(utils.ErrorAuth)
			c.Abort()
			return
		}

		// Parse and validate the token
		tokenString := parts[1]
		logrus.Infof("token: %v", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logrus.Errorf("Invalid Signing Method")
				return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
			}
			jwtSecretKey := config.Init().JWTSecret
			logrus.Infof("jwt secret key from token: %v", jwtSecretKey)

			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			c.Error(utils.ErrorAuth)
			c.Abort()
			return
		}

		// Token is valid, extract claims if needed
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set(user_id_key, claims[user_id_key_claim])
		}

		if isUserExist := userUseCase.CheckUserExist(c.GetString(user_id_key)); isUserExist == false {
			c.Error(utils.ErrorAuth)
			c.Abort()
			return
		}

		c.Next() // Proceed to the next handler
	}
}
