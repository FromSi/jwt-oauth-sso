package responses

import (
	"encoding/json"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSuccessDevicesResponse(t *testing.T) {
	var devices []repositories.Device

	device := repositories.NewGormDevice()

	device.SetUUID("1")
	device.SetUserUUID("2")
	device.SetUserAgent("3")
	device.SetIp("4")
	device.SetRefreshToken("5")
	device.SetExpiresAt(6)
	device.SetCreatedAt(7)
	device.SetUpdatedAt(8)

	devices = append(devices, device)

	response := NewSuccessDevicesResponse(devices)

	assert.NotEmpty(t, response)

	responseToJson, err := json.Marshal(response)

	assert.NoError(t, err)

	expected := `{"data":[{"uuid":"1","userUUID":"2","userAgent":"3","ip":"4","expiresAt":6,"createdAt":7,"updatedAt":8}]}`

	assert.Equal(t, string(responseToJson), expected)
}
