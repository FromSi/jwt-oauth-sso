package responses

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorConflictResponse(t *testing.T) {
	response := NewErrorConflictResponse("value")

	assert.NotNil(t, response)

	responseToJson, err := json.Marshal(response)

	assert.Nil(t, err)

	expected := `{"message":"value"}`

	assert.Equal(t, string(responseToJson), expected)
}
