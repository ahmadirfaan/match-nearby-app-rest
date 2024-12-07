// errors.go
package middleware

import "errors"

var (
	ErrorAuth      = errors.New("AuthorizationError")
	ErrorNotFound  = errors.New("NotFoundError")
	ErrorForbidden = errors.New("ForbiddenError")
)
