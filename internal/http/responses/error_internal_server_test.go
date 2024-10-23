package responses

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorInternalServerResponse(t *testing.T) {
	response := NewErrorInternalServerResponse("value")

	assert.NotNil(t, response)

	responseToJson, err := json.Marshal(response)

	assert.Nil(t, err)

	expected := `{"message":"value"}`

	assert.Equal(t, string(responseToJson), expected)
}
