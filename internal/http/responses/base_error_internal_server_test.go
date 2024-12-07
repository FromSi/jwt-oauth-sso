package responses

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseErrorInternalServerResponse(t *testing.T) {
	response := NewBaseErrorInternalServerResponse()

	assert.NotNil(t, response)
}

func TestNewBaseErrorInternalServerResponse_Make(t *testing.T) {
	responseConstructor := NewBaseErrorInternalServerResponse()

	assert.NotNil(t, responseConstructor)

	response := responseConstructor.Make(errors.New("1"))

	assert.NotEmpty(t, response)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"message":"1"}`

	assert.Equal(t, expected, string(responseToJson))

	response = responseConstructor.Make(errors.New("2"))

	assert.NotEmpty(t, response)

	responseToJson, err = json.Marshal(response)

	assert.NoError(t, err)

	expected = `{"message":"2"}`

	assert.Equal(t, expected, string(responseToJson))
}
