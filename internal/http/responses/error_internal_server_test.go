package responses

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorInternalServerResponse(t *testing.T) {
	response := NewErrorInternalServerResponse(errors.New("error"))

	assert.NotEmpty(t, response)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"message":"error"}`

	assert.Equal(t, string(responseToJson), expected)
}
