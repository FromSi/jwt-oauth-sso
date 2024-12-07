package responses

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseErrorConflictResponse(t *testing.T) {
	response := NewBaseErrorConflictResponse()

	assert.NotNil(t, response)
}

func TestNewBaseErrorConflictResponse_Make(t *testing.T) {
	responseConstructor := NewBaseErrorConflictResponse()

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
