package responses

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorBadRequestResponse(t *testing.T) {
	data := make(map[string][]string)
	data["field"] = append(data["field"], "1")
	data["field"] = append(data["field"], "2")

	response := NewErrorBadRequestResponse(data)

	assert.NotNil(t, response)

	assert.Equal(t, response.Errors["field"][0], "1")
	assert.Equal(t, response.Errors["field"][1], "2")

	responseToJson, err := json.Marshal(response)

	assert.Nil(t, err)

	expected := `{"errors":{"field":["1","2"]}}`

	assert.Equal(t, string(responseToJson), expected)
}

func Test_NewErrorBadRequestResponseByError(t *testing.T) {
	err := errors.New("some other error")
	response := NewErrorBadRequestResponseByError(err)

	assert.Contains(t, response.Errors, "error")
	assert.Equal(t, response.Errors["error"][0], "some other error")

	err = validator.New().Struct(struct {
		Email string `validate:"required,email"`
	}{})

	var validationErrors validator.ValidationErrors

	assert.ErrorAs(t, err, &validationErrors)

	response = NewErrorBadRequestResponseByError(err)

	assert.Contains(t, response.Errors, "email")
	assert.Equal(t, response.Errors["email"][0], "validation failed on required")
}
