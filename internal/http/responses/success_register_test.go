package responses

import (
	"encoding/json"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSuccessRegisterResponse(t *testing.T) {
	config := configs.NewBaseConfig()
	gormDevice := repositories.NewGormDevice()
	response, err := NewSuccessRegisterResponse(config, gormDevice)

	assert.Nil(t, err)
	assert.NotNil(t, response)

	response.Data.AccessToken = "1"
	response.Data.RefreshToken = "2"
	response.Data.AccessExpiresIn = 3
	response.Data.RefreshExpiresIn = 4

	responseToJson, err := json.Marshal(response)

	assert.Nil(t, err)

	expected := `{"data":{"authType":"bearer","accessToken":"1","refreshToken":"2","accessExpiresIn":3,"refreshExpiresIn":4}}`

	assert.Equal(t, string(responseToJson), expected)
}
