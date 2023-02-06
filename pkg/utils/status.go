package utils

import (
	"net/http"

	custom_error "github.com/fauziagustian/delos-aqua/pkg/customError"
)

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*custom_error.NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*custom_error.UserAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*custom_error.IncorrectCredentialsError); ok {
		statusCode = http.StatusUnauthorized
	} else if _, ok := err.(*custom_error.UserNotFoundError); ok {
		statusCode = http.StatusBadRequest
	}
	return statusCode
}
