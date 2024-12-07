package responses

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseErrorBadRequestResponse(t *testing.T) {
	response := NewBaseErrorBadRequestResponse()

	assert.NotNil(t, response)
}

func TestNewBaseErrorBadRequestResponse_Make(t *testing.T) {
	responseConstructor := NewBaseErrorBadRequestResponse()

	assert.NotNil(t, responseConstructor)

	err := validator.New().Struct(struct {
		Email string `validate:"required,email"`
	}{})

	var validationErrors validator.ValidationErrors

	assert.ErrorAs(t, err, &validationErrors)

	response := responseConstructor.Make(err)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"errors":{"email":["validation failed on required"]}}`

	assert.Equal(t, expected, string(responseToJson))

	err = errors.New("error")

	response = responseConstructor.Make(err)

	responseToJson, err = json.Marshal(response)

	assert.NoError(t, err)

	expected = `{"errors":{"error":["error"]}}`

	assert.Equal(t, expected, string(responseToJson))
}
