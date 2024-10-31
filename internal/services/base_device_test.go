package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/internal/mocks/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseDeviceService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	baseDeviceService := NewBaseDeviceService(mockDeviceRepository)

	assert.NotNil(t, baseDeviceService)
}

func TestBaseDeviceService_GenerateUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	baseDeviceService := NewBaseDeviceService(mockDeviceRepository)

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
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	baseDeviceService := NewBaseDeviceService(mockDeviceRepository)

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

func TestBaseDeviceService_GetNewDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	device := repositories.NewGormDevice()

	device.SetUserUUID("1")
	device.SetIp("1")
	device.SetUserAgent("1")
	device.SetRefreshToken("1")

	mockDeviceRepository.EXPECT().GetDeviceByUserUUIDAndIpAndUserAgent("1", "1", "1").Return(device).AnyTimes()
	mockDeviceRepository.EXPECT().GetDeviceByUserUUIDAndIpAndUserAgent("2", "2", "2").Return(nil).AnyTimes()

	baseDeviceService := NewBaseDeviceService(mockDeviceRepository)

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
			device := baseDeviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(config, tt.userUUID, tt.ip, tt.agent)

			assert.NotNil(t, device)

			assert.Equal(t, tt.userUUID, device.GetUserUUID())
			assert.Equal(t, tt.ip, device.GetIp())
			assert.Equal(t, tt.agent, device.GetUserAgent())
			assert.NotEmpty(t, device.GetRefreshToken())
		})
	}
}
