package responses

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ErrorBadRequestResponse struct {
	Errors map[string][]string `json:"errors"`
}

func NewErrorBadRequestResponse(errors map[string][]string) *ErrorBadRequestResponse {
	return &ErrorBadRequestResponse{
		Errors: errors,
	}
}

func NewErrorBadRequestResponseByError(err error) *ErrorBadRequestResponse {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		errorsMap := make(map[string][]string)

		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			field = strings.ToLower(field[:1]) + field[1:]
			tag := fieldError.Tag()
			errorsMap[field] = append(errorsMap[field], "validation failed on "+tag)
		}

		return &ErrorBadRequestResponse{
			Errors: errorsMap,
		}
	}

	return &ErrorBadRequestResponse{
		Errors: map[string][]string{"error": {err.Error()}},
	}
}
