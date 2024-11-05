package responses

import (
	"encoding/json"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSuccessDevicesResponse(t *testing.T) {
	var devices []repositories.Device

	gormDevice := repositories.NewGormDevice()

	gormDevice.SetUUID("1")
	gormDevice.SetUserUUID("2")
	gormDevice.SetUserAgent("3")
	gormDevice.SetIp("4")
	gormDevice.SetRefreshToken("5")
	gormDevice.SetExpiredAt(6)
	gormDevice.SetCreatedAt(7)
	gormDevice.SetUpdatedAt(8)

	devices = append(devices, gormDevice)

	response := NewSuccessDevicesResponse(devices)

	assert.NotNil(t, response)

	responseToJson, err := json.Marshal(response)

	assert.Nil(t, err)

	expected := `{"data":[{"uuid":"1","userUUID":"2","userAgent":"3","ip":"4","expiredAt":6,"createdAt":7,"updatedAt":8}]}`

	assert.Equal(t, string(responseToJson), expected)
}
