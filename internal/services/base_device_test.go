package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseDeviceService(t *testing.T) {
	mockDeviceRepository := MockDeviceRepository{}

	baseDeviceService := NewBaseDeviceService(&mockDeviceRepository)

	assert.NotNil(t, baseDeviceService)
}

func TestBaseDeviceService_GenerateUUID(t *testing.T) {
	mockDeviceRepository := MockDeviceRepository{}

	baseDeviceService := NewBaseDeviceService(&mockDeviceRepository)

	uuidOne := baseDeviceService.GenerateUUID()
	uuidTwo := baseDeviceService.GenerateUUID()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.Nil(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.Nil(t, err)
}

func TestBaseDeviceService_GenerateRefreshToken(t *testing.T) {
	mockDeviceRepository := MockDeviceRepository{}

	baseDeviceService := NewBaseDeviceService(&mockDeviceRepository)

	uuidOne := baseDeviceService.GenerateRefreshToken()
	uuidTwo := baseDeviceService.GenerateRefreshToken()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.Nil(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.Nil(t, err)
}

func TestBaseDeviceService_GetDeviceByUserUUIDAndIpAndAgent(t *testing.T) {
	mockDeviceRepository := MockDeviceRepository{}

	device := repositories.NewGormDevice()

	device.SetUserUUID("1")
	device.SetIp("1")
	device.SetAgent("1")

	mockDeviceRepository.On("GetDeviceByUserUUIDAndIpAndAgent", "1", "1", "1").Return(device)
	mockDeviceRepository.On("GetDeviceByUserUUIDAndIpAndAgent", "2", "2", "2").Return(nil)

	baseDeviceService := NewBaseDeviceService(&mockDeviceRepository)

	tests := []struct {
		name     string
		userUUID string
		ip       string
		agent    string
	}{
		{
			name:     "Found device",
			userUUID: "1",
			ip:       "1",
			agent:    "1",
		},
		{
			name:     "Not found device, but create one",
			userUUID: "2",
			ip:       "2",
			agent:    "2",
		},
	}

	config := configs.NewBaseConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device := baseDeviceService.GetDeviceByUserUUIDAndIpAndAgent(config, tt.userUUID, tt.ip, tt.agent)

			assert.NotNil(t, device)

			assert.Equal(t, tt.userUUID, device.GetUserUUID())
			assert.Equal(t, tt.ip, device.GetIp())
			assert.Equal(t, tt.agent, device.GetAgent())
		})
	}
}
