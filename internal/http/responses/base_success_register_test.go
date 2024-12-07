package responses

import (
	"encoding/json"
	"errors"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseSuccessRegisterResponse(t *testing.T) {
	response := NewBaseSuccessRegisterResponse()

	assert.NotNil(t, response)
}

func TestNewBaseSuccessRegisterResponse_Make(t *testing.T) {
	responseConstructor := NewBaseSuccessRegisterResponse()

	assert.NotNil(t, responseConstructor)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockDevice := repositories_mocks.NewMockDevice(mockController)

	tests := []struct {
		name             string
		valueInt         int
		valueString      string
		expected         string
		errorAccessToken error
		errorDevice      error
	}{
		{
			name:             "Value one",
			valueInt:         1,
			valueString:      "1",
			expected:         `{"data":{"authType":"bearer","accessToken":"1","refreshToken":"1","accessExpiresIn":1,"refreshExpiresIn":1}}`,
			errorAccessToken: nil,
			errorDevice:      nil,
		},
		{
			name:             "Value two",
			valueInt:         2,
			valueString:      "2",
			expected:         `{"data":{"authType":"bearer","accessToken":"2","refreshToken":"2","accessExpiresIn":2,"refreshExpiresIn":2}}`,
			errorAccessToken: nil,
			errorDevice:      nil,
		},
		{
			name:             "Error access token",
			valueInt:         0,
			valueString:      "0",
			expected:         ``,
			errorAccessToken: errors.New("error"),
			errorDevice:      nil,
		},
		{
			name:             "Error device",
			valueInt:         0,
			valueString:      "0",
			expected:         ``,
			errorAccessToken: nil,
			errorDevice:      errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.errorAccessToken == nil && tt.errorDevice == nil {
				mockAccessToken.EXPECT().GetExpirationTime().Return(tt.valueInt)
				mockDevice.EXPECT().GetRefreshToken().Return(tt.valueString)
				mockDevice.EXPECT().GetExpiresAt().Return(tt.valueInt)
			}

			if tt.errorDevice == nil {
				mockAccessToken.EXPECT().ToString().Return(tt.valueString, tt.errorAccessToken)
			}

			mockDevice.EXPECT().GenerateAccessToken().Return(mockAccessToken, tt.errorDevice)

			response, err := responseConstructor.Make(mockDevice)

			if tt.errorAccessToken != nil || tt.errorDevice != nil {
				assert.Error(t, err)
				assert.Nil(t, response)

				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, response)

			responseToJson, err := json.Marshal(response)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(responseToJson))
		})
	}
}
