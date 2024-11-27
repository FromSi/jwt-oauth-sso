package responses

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorInternalServerResponse(t *testing.T) {
	response := NewErrorInternalServerResponse(errors.New("1"))

	assert.NotEmpty(t, response)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"message":"1"}`

	assert.Equal(t, expected, string(responseToJson))

	response = NewErrorInternalServerResponse(errors.New("2"))

	assert.NotEmpty(t, response)

	responseToJson, err = json.Marshal(response)

	assert.NoError(t, err)

	expected = `{"message":"2"}`

	assert.Equal(t, expected, string(responseToJson))
}
