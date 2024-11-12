package responses

import (
	"encoding/json"
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewSuccessLoginResponse(t *testing.T) {
	config := configs.NewBaseConfig()
	device := repositories.NewGormDevice()

	response, err := NewSuccessLoginResponse(config, device)

	assert.NoError(t, err)
	assert.NotEmpty(t, response)

	response.Data.AccessToken = "1"
	response.Data.RefreshToken = "2"
	response.Data.AccessExpiresIn = 3
	response.Data.RefreshExpiresIn = 4

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"data":{"authType":"bearer","accessToken":"1","refreshToken":"2","accessExpiresIn":3,"refreshExpiresIn":4}}`

	assert.Equal(t, string(responseToJson), expected)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	deviceMock := repositories_mocks.NewMockDevice(mockController)

	deviceMock.
		EXPECT().
		GenerateAccessToken(gomock.Any()).
		Return(nil, errors.New("error")).
		AnyTimes()

	response, err = NewSuccessLoginResponse(config, deviceMock)

	assert.Error(t, err)
	assert.Empty(t, response)
}
