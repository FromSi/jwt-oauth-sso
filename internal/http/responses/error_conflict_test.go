package responses

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorConflictResponse(t *testing.T) {
	response := NewErrorConflictResponse(errors.New("error"))

	assert.NotNil(t, response)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"message":"error"}`

	assert.Equal(t, string(responseToJson), expected)
}
