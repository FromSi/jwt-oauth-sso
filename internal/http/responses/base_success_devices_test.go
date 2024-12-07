package responses

import (
	"encoding/json"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseSuccessDevicesResponse(t *testing.T) {
	response := NewBaseSuccessDevicesResponse()

	assert.NotNil(t, response)
}

func TestNewBaseSuccessDevicesResponse_Make(t *testing.T) {
	responseConstructor := NewBaseSuccessDevicesResponse()

	assert.NotNil(t, responseConstructor)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDevice := repositories_mocks.NewMockDevice(mockController)

	tests := []struct {
		name        string
		valueInt    int
		valueString string
		expected    string
	}{
		{
			name:        "Value one",
			valueInt:    1,
			valueString: "1",
			expected:    `{"data":[{"uuid":"1","userUUID":"1","userAgent":"1","ip":"1","issuedAt":1,"expiresAt":1,"createdAt":1,"updatedAt":1}]}`,
		},
		{
			name:        "Value two",
			valueInt:    2,
			valueString: "2",
			expected:    `{"data":[{"uuid":"2","userUUID":"2","userAgent":"2","ip":"2","issuedAt":2,"expiresAt":2,"createdAt":2,"updatedAt":2}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var devices []repositories.Device

			mockDevice.EXPECT().GetUUID().Return(tt.valueString)
			mockDevice.EXPECT().GetUserUUID().Return(tt.valueString)
			mockDevice.EXPECT().GetUserAgent().Return(tt.valueString)
			mockDevice.EXPECT().GetIp().Return(tt.valueString)
			mockDevice.EXPECT().GetIssuedAt().Return(tt.valueInt)
			mockDevice.EXPECT().GetExpiresAt().Return(tt.valueInt)
			mockDevice.EXPECT().GetCreatedAt().Return(tt.valueInt)
			mockDevice.EXPECT().GetUpdatedAt().Return(tt.valueInt)

			devices = append(devices, mockDevice)

			response := responseConstructor.Make(devices)

			assert.NotEmpty(t, response)

			responseToJson, err := json.Marshal(response)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(responseToJson))
		})
	}
}
