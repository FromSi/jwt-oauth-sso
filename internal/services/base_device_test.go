package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
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
	mockDeviceBuilder := repositories_mocks.NewMockDeviceBuilder(mockController)

	baseDeviceService := NewBaseDeviceService(
		config,
		mockDeviceRepository,
		mockDeviceBuilder,
	)

	assert.NotEmpty(t, baseDeviceService)
}

func TestBaseDeviceService_GenerateUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceBuilder := repositories_mocks.NewMockDeviceBuilder(mockController)

	baseDeviceService := NewBaseDeviceService(
		config,
		mockDeviceRepository,
		mockDeviceBuilder,
	)

	uuidOne := baseDeviceService.GenerateUUID()
	uuidTwo := baseDeviceService.GenerateUUID()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.NoError(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.NoError(t, err)
}

func TestBaseDeviceService_GenerateRefreshToken(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceBuilder := repositories_mocks.NewMockDeviceBuilder(mockController)

	baseDeviceService := NewBaseDeviceService(
		config,
		mockDeviceRepository,
		mockDeviceBuilder,
	)

	uuidOne := baseDeviceService.GenerateRefreshToken()
	uuidTwo := baseDeviceService.GenerateRefreshToken()

	assert.NotEmpty(t, uuidOne)
	assert.NotEmpty(t, uuidTwo)

	assert.NotEqual(t, uuidOne, uuidTwo)

	_, err := uuid.Parse(uuidOne)

	assert.NoError(t, err)

	_, err = uuid.Parse(uuidTwo)

	assert.NoError(t, err)
}

func TestBaseDeviceService_GetOldDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceBuilder := repositories_mocks.NewMockDeviceBuilder(mockController)
	mockDevice := repositories_mocks.NewMockDevice(mockController)

	mockDeviceRepository.
		EXPECT().
		GetDeviceByUserUUIDAndIpAndUserAgent(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(mockDevice).
		AnyTimes()

	baseDeviceService := NewBaseDeviceService(
		config,
		mockDeviceRepository,
		mockDeviceBuilder,
	)

	device := baseDeviceService.GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "1", "1")

	assert.NotEmpty(t, device)
}

func TestBaseDeviceService_GetNewDeviceByUserUUIDAndIpAndUserAgent(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceBuilder := repositories_mocks.NewMockDeviceBuilder(mockController)
	mockDevice := repositories_mocks.NewMockDevice(mockController)

	mockDeviceBuilder.EXPECT().New().Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetUUID(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetUserUUID(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetIp(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetUserAgent(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetCreatedAt(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetUpdatedAt(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().Build().Return(mockDevice, nil).AnyTimes()

	baseDeviceService := NewBaseDeviceService(
		config,
		mockDeviceRepository,
		mockDeviceBuilder,
	)

	device, err := baseDeviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "1", "1")

	assert.NoError(t, err)
	assert.NotEmpty(t, device)
}

func TestBaseDeviceService_GetNewRefreshDetailsByDevice(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	config := configs.NewBaseConfig()
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceBuilder := repositories_mocks.NewMockDeviceBuilder(mockController)
	mockDeviceOne := repositories_mocks.NewMockDevice(mockController)
	mockDeviceTwo := repositories_mocks.NewMockDevice(mockController)

	mockDeviceOne.EXPECT().SetRefreshToken(gomock.Any()).Return().AnyTimes()
	mockDeviceOne.EXPECT().SetUpdatedAt(gomock.Any()).Return().AnyTimes()
	mockDeviceOne.EXPECT().SetExpiresAt(gomock.Any()).Return().AnyTimes()

	mockDeviceBuilder.EXPECT().NewFromDevice(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetRefreshToken(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetExpiresAt(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().SetUpdatedAt(gomock.Any()).Return(mockDeviceBuilder).AnyTimes()
	mockDeviceBuilder.EXPECT().Build().Return(mockDeviceTwo, nil)

	baseDeviceService := NewBaseDeviceService(
		config,
		mockDeviceRepository,
		mockDeviceBuilder,
	)

	deviceUpdated, err := baseDeviceService.GetNewRefreshDetailsByDevice(mockDeviceOne)

	assert.NoError(t, err)
	assert.NotEmpty(t, deviceUpdated)
	assert.NotSame(t, mockDeviceOne, deviceUpdated)

	mockDeviceBuilder.EXPECT().Build().Return(nil, errors.New("error"))

	deviceUpdated, err = baseDeviceService.GetNewRefreshDetailsByDevice(mockDeviceOne)

	assert.Error(t, err)
	assert.Empty(t, deviceUpdated)
}
