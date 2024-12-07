package responses

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type BaseErrorBadRequestResponse struct {
	Errors map[string][]string `json:"errors"`
}

func NewBaseErrorBadRequestResponse() *BaseErrorBadRequestResponse {
	return &BaseErrorBadRequestResponse{}
}

func (receiver BaseErrorBadRequestResponse) Make(
	err error,
) ErrorBadRequestResponse {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		errorsMap := make(map[string][]string)

		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			field = strings.ToLower(field[:1]) + field[1:]
			tag := fieldError.Tag()
			errorsMap[field] = append(errorsMap[field], "validation failed on "+tag)
		}

		return &BaseErrorBadRequestResponse{
			Errors: errorsMap,
		}
	}

	return &BaseErrorBadRequestResponse{
		Errors: map[string][]string{"error": {err.Error()}},
	}
}
