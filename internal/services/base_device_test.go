package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseDeviceService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	baseDeviceService := NewBaseDeviceService(config, mockDeviceRepository)

	assert.NotNil(t, baseDeviceService)
}

func TestBaseDeviceService_GenerateUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	baseDeviceService := NewBaseDeviceService(config, mockDeviceRepository)

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

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	baseDeviceService := NewBaseDeviceService(config, mockDeviceRepository)

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

func TestBaseDeviceService_GetOldDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	deviceOne := repositories.NewGormDevice()

	deviceOne.SetUserUUID("1")
	deviceOne.SetIp("1")
	deviceOne.SetUserAgent("1")

	mockDeviceRepository.
		EXPECT().
		GetDeviceByUserUUIDAndIpAndUserAgent(
			deviceOne.GetUserUUID(),
			deviceOne.GetIp(),
			deviceOne.GetUserAgent(),
		).
		Return(deviceOne).
		AnyTimes()

	mockDeviceRepository.
		EXPECT().
		GetDeviceByUserUUIDAndIpAndUserAgent(
			"0",
			"0",
			"0",
		).
		Return(nil).
		AnyTimes()

	baseDeviceService := NewBaseDeviceService(config, mockDeviceRepository)

	device := baseDeviceService.
		GetOldDeviceByUserUUIDAndIpAndUserAgent(
			deviceOne.GetUserUUID(),
			deviceOne.GetIp(),
			deviceOne.GetUserAgent(),
		)

	assert.NotNil(t, device)

	assert.Equal(t, deviceOne.GetUserUUID(), device.GetUserUUID())
	assert.Equal(t, deviceOne.GetIp(), device.GetIp())
	assert.Equal(t, deviceOne.GetUserAgent(), device.GetUserAgent())

	device = baseDeviceService.
		GetOldDeviceByUserUUIDAndIpAndUserAgent(
			"0",
			"0",
			"0",
		)

	assert.Nil(t, device)
}

func TestBaseDeviceService_GetNewDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)

	deviceOne := repositories.NewGormDevice()

	deviceOne.SetUserUUID("1")
	deviceOne.SetIp("1")
	deviceOne.SetUserAgent("1")

	deviceTwo := repositories.NewGormDevice()

	deviceTwo.SetUserUUID("2")
	deviceTwo.SetIp("2")
	deviceTwo.SetUserAgent("2")

	baseDeviceService := NewBaseDeviceService(config, mockDeviceRepository)

	device := baseDeviceService.
		GetNewDeviceByUserUUIDAndIpAndUserAgent(
			deviceOne.GetUserUUID(),
			deviceOne.GetIp(),
			deviceOne.GetUserAgent(),
		)

	assert.NotNil(t, device)

	assert.Equal(t, deviceOne.GetUserUUID(), device.GetUserUUID())
	assert.Equal(t, deviceOne.GetIp(), device.GetIp())
	assert.Equal(t, deviceOne.GetUserAgent(), device.GetUserAgent())

	device = baseDeviceService.
		GetNewDeviceByUserUUIDAndIpAndUserAgent(
			deviceTwo.GetUserUUID(),
			deviceTwo.GetIp(),
			deviceTwo.GetUserAgent(),
		)

	assert.NotNil(t, device)

	assert.Equal(t, deviceTwo.GetUserUUID(), device.GetUserUUID())
	assert.Equal(t, deviceTwo.GetIp(), device.GetIp())
	assert.Equal(t, deviceTwo.GetUserAgent(), device.GetUserAgent())
}
