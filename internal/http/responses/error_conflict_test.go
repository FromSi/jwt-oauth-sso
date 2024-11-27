package responses

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorConflictResponse(t *testing.T) {
	response := NewErrorConflictResponse(errors.New("1"))

	assert.NotEmpty(t, response)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"message":"1"}`

	assert.Equal(t, expected, string(responseToJson))

	response = NewErrorConflictResponse(errors.New("2"))

	assert.NotEmpty(t, response)

	responseToJson, err = json.Marshal(response)

	assert.NoError(t, err)

	expected = `{"message":"2"}`

	assert.Equal(t, expected, string(responseToJson))
}
