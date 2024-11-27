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

	assert.NotEmpty(t, response)

	assert.Equal(t, response.Errors["field"][0], "1")
	assert.Equal(t, response.Errors["field"][1], "2")

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"errors":{"field":["1","2"]}}`

	assert.Equal(t, expected, string(responseToJson))

	response = NewErrorBadRequestResponseByError(errors.New("error"))

	assert.Contains(t, response.Errors, "error")
	assert.Equal(t, response.Errors["error"][0], "error")

	err = validator.New().Struct(struct {
		Email string `validate:"required,email"`
	}{})

	var validationErrors validator.ValidationErrors

	assert.ErrorAs(t, err, &validationErrors)

	response = NewErrorBadRequestResponseByError(err)

	assert.Contains(t, response.Errors, "email")
	assert.Equal(t, "validation failed on required", response.Errors["email"][0])
}
