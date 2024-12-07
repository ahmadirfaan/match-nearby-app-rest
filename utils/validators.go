package utils

import (
	"github.com/go-playground/validator"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

