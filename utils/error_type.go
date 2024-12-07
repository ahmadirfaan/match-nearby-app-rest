// errors.go
package utils

import "errors"

var (
	ErrorAuth       = errors.New("AuthorizationError")
	ErrorNotFound   = errors.New("NotFoundError")
	ErrorForbidden  = errors.New("ForbiddenError")
	ErrorBadRequest = errors.New("BadRequestError")
	ErrorValidator  = errors.New("ErrorValidator")
)
