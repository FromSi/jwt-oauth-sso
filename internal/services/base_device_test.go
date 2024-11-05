package services

import (
	"errors"
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

func TestBaseDeviceService_GetDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
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
	mockDeviceRepository.EXPECT().UpdateDevice(gomock.Any()).Return(nil).AnyTimes()

	baseDeviceService := NewBaseDeviceService(mockDeviceRepository)

	tests := []struct {
		name     string
		userUUID string
		ip       string
		agent    string
		isEmpty  bool
	}{
		{
			name:     "Found device",
			userUUID: "1",
			ip:       "1",
			agent:    "1",
			isEmpty:  false,
		},
		{
			name:     "Not found device, but create one",
			userUUID: "2",
			ip:       "2",
			agent:    "2",
			isEmpty:  true,
		},
	}

	config := configs.NewBaseConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device, err := baseDeviceService.GetDeviceByUserUUIDAndIpAndUserAgent(config, tt.userUUID, tt.ip, tt.agent)

			assert.Nil(t, err)

			if tt.isEmpty {
				assert.Nil(t, device)

				return
			}

			assert.NotNil(t, device)

			assert.Equal(t, tt.userUUID, device.GetUserUUID())
			assert.Equal(t, tt.ip, device.GetIp())
			assert.Equal(t, tt.agent, device.GetUserAgent())
			assert.NotEmpty(t, device.GetRefreshToken())
		})
	}
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
	mockDeviceRepository.EXPECT().CreateDevice(gomock.Any()).Return(nil).AnyTimes()

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
			device, err := baseDeviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(config, tt.userUUID, tt.ip, tt.agent)

			assert.Nil(t, err)
			assert.NotNil(t, device)

			assert.Equal(t, tt.userUUID, device.GetUserUUID())
			assert.Equal(t, tt.ip, device.GetIp())
			assert.Equal(t, tt.agent, device.GetUserAgent())
			assert.NotEmpty(t, device.GetRefreshToken())
		})
	}
}

func TestBaseDeviceService_ResetDevice(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	deviceOne := repositories.NewGormDevice()
	deviceTwo := repositories.NewGormDevice()

	deviceOne.SetUUID("1")
	deviceOne.SetRefreshToken("1")
	deviceTwo.SetUUID("2")
	deviceTwo.SetRefreshToken("2")

	mockDeviceRepository.EXPECT().UpdateDevice(gomock.Cond(func(device repositories.Device) bool {
		return device.GetUUID() == deviceTwo.GetUUID()
	})).Return(errors.New("error")).AnyTimes()

	mockDeviceRepository.EXPECT().UpdateDevice(gomock.Cond(func(device repositories.Device) bool {
		return device.GetRefreshToken() != deviceOne.GetRefreshToken()
	})).Return(nil).AnyTimes()

	baseDeviceService := NewBaseDeviceService(mockDeviceRepository)

	config := configs.NewBaseConfig()

	device, err := baseDeviceService.ResetDevice(config, deviceTwo)

	assert.NotNil(t, err)
	assert.Nil(t, device)

	device, err = baseDeviceService.ResetDevice(config, deviceOne)

	assert.Nil(t, err)
	assert.NotNil(t, device)
	assert.NotEqual(t, device.GetRefreshToken(), deviceOne.GetRefreshToken())
}
