package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func FormatValidationError(err error) []ErrorMsg {
	errorMessages := []ErrorMsg{}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		e := strings.Split(err.Error(), ".")[1]
		errors := strings.Split(e, " of type ")
		message := fmt.Sprintf("%s must be %s", errors[0], errors[1])
		errorMessages = append(errorMessages, ErrorMsg{Field: errors[0], Message: message})
		return errorMessages
	}

	if !errors.Is(err, io.EOF) {
		for _, e := range err.(validator.ValidationErrors) {
			var message string
			if e.Param() != "" {
				message = ToSnakeCase(e.Field()) + " " + e.Tag() + " " + e.Param()
			} else {
				message = ToSnakeCase(e.Field()) + " " + e.Tag()
			}
			errorMessages = append(errorMessages, ErrorMsg{Field: ToSnakeCase(e.Field()), Message: message})
		}
		return errorMessages
	}
	return errorMessages
}
