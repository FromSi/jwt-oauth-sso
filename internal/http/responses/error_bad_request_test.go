package responses

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorBadRequestResponse(t *testing.T) {
	data := make(map[string][]string)
	data["field"] = append(data["field"], "valueOne")
	data["field"] = append(data["field"], "valueTwo")

	response := NewErrorBadRequestResponse(data)

	assert.NotNil(t, response)

	assert.Equal(t, response.Errors["field"][0], "valueOne")
	assert.Equal(t, response.Errors["field"][1], "valueTwo")

	responseToJson, err := json.Marshal(response)

	assert.Nil(t, err)

	expected := `{"errors":{"field":["valueOne","valueTwo"]}}`

	assert.Equal(t, string(responseToJson), expected)
}
