package web

import "net/http"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AuthError() ErrorResponse {
	return newErrorResponse(http.StatusUnauthorized, "Authentication failed. Please check your credentials.")
}

func NotFoundError() ErrorResponse {
	return newErrorResponse(http.StatusNotFound, "The requested resource was not found.")
}

func ForbiddenError() ErrorResponse {
	return newErrorResponse(http.StatusForbidden, "You do not have permission to access this resource.")
}

func InternalServiceError() ErrorResponse {
	return newErrorResponse(http.StatusInternalServerError, "Internal Server Error")
}

func newErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}
