// middleware.go
package middleware

import (
	"errors"

	"github.com/ahmadirfaan/match-nearby-app-rest/models/web"
	"github.com/ahmadirfaan/match-nearby-app-rest/utils"
	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware for handling errors globally
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process the request
		c.Next()

		// Check for errors after processing the request
		if len(c.Errors) > 0 {
			// If error is found, log and handle it
			err := c.Errors.Last()

			var resp web.ErrorResponse
			switch {
			case errors.Is(err.Err, utils.ErrorAuth):
				resp = web.AuthError()
			case errors.Is(err.Err, utils.ErrorNotFound):
				resp = web.NotFoundError()
			case errors.Is(err.Err, utils.ErrorForbidden):
				resp = web.ForbiddenError()
			case errors.Is(err.Err, utils.ErrorBadRequest) || errors.Is(err.Err, utils.ErrorValidator):
				resp = web.BadRequestError()
			default:
				resp = web.InternalServiceError()
			}

			// Respond with the custom error
			c.JSON(resp.Code, resp)
		}
	}
}

func NoRouteHandler(c *gin.Context) {
	// Raise the ErrNotFound error
	c.Error(utils.ErrorNotFound)
}
