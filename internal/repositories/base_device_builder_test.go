package repositories

import (
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewBaseDeviceBuilder(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)
}

func TestBaseDeviceBuilder_New(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilderOne := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilderOne)

	baseDeviceBuilderOne.SetUUID("1")

	baseDeviceBuilderTwo := baseDeviceBuilderOne.New()

	assert.NotNil(t, baseDeviceBuilderTwo)
	assert.NotEqual(t, baseDeviceBuilderOne, baseDeviceBuilderTwo)
}

func TestBaseDeviceBuilder_NewFromDevice(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilderOne := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilderOne)

	baseDeviceBuilderOne.SetUUID("1")

	user := GormDevice{UUID: "2"}

	baseDeviceBuilderTwo := baseDeviceBuilderOne.NewFromDevice(&user)

	assert.NotNil(t, baseDeviceBuilderTwo)
	assert.NotEqual(t, baseDeviceBuilderOne, baseDeviceBuilderTwo)
}

func TestBaseDeviceBuilder_Build(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	baseDeviceBuilder.SetUUID("1")

	device, err := baseDeviceBuilder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, device)

	assert.Equal(t, baseDeviceBuilder.device.GetUUID(), device.GetUUID())
}

func TestBaseDeviceBuilder_BuildToGorm(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	baseDeviceBuilder.SetUUID("1")

	device, err := baseDeviceBuilder.BuildToGorm()

	assert.NoError(t, err)
	assert.NotNil(t, device)

	assert.Equal(t, baseDeviceBuilder.device.GetUUID(), device.GetUUID())

	baseDeviceBuilder.SetUUID("")

	device, err = baseDeviceBuilder.BuildToGorm()

	assert.Error(t, err)
	assert.Nil(t, device)
}

func TestBaseDeviceBuilder_SetUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetUUID("1")

	assert.Equal(t, "1", baseDeviceBuilder.device.GetUUID())

	baseDeviceBuilder.SetUUID("2")

	assert.Equal(t, "2", baseDeviceBuilder.device.GetUUID())
}

func TestBaseDeviceBuilder_SetUserUUID(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetUserUUID("1")

	assert.Equal(t, "1", baseDeviceBuilder.device.GetUserUUID())

	baseDeviceBuilder.SetUserUUID("2")

	assert.Equal(t, "2", baseDeviceBuilder.device.GetUserUUID())
}

func TestBaseDeviceBuilder_SetUserAgent(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetUserAgent("1")

	assert.Equal(t, "1", baseDeviceBuilder.device.GetUserAgent())

	baseDeviceBuilder.SetUserAgent("2")

	assert.Equal(t, "2", baseDeviceBuilder.device.GetUserAgent())
}

func TestBaseDeviceBuilder_SetIp(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetIp("1")

	assert.Equal(t, "1", baseDeviceBuilder.device.GetIp())

	baseDeviceBuilder.SetIp("2")

	assert.Equal(t, "2", baseDeviceBuilder.device.GetIp())
}

func TestBaseDeviceBuilder_SetRefreshToken(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetRefreshToken("1")

	assert.Equal(t, "1", baseDeviceBuilder.device.GetRefreshToken())

	baseDeviceBuilder.SetRefreshToken("2")

	assert.Equal(t, "2", baseDeviceBuilder.device.GetRefreshToken())
}

func TestBaseDeviceBuilder_SetExpiresAt(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetExpiresAt(1)

	assert.Equal(t, 1, baseDeviceBuilder.device.GetExpiresAt())

	baseDeviceBuilder.SetExpiresAt(2)

	assert.Equal(t, 2, baseDeviceBuilder.device.GetExpiresAt())
}

func TestBaseDeviceBuilder_SetCreatedAt(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetCreatedAt(1)

	assert.Equal(t, 1, baseDeviceBuilder.device.GetCreatedAt())

	baseDeviceBuilder.SetCreatedAt(2)

	assert.Equal(t, 2, baseDeviceBuilder.device.GetCreatedAt())
}

func TestBaseDeviceBuilder_SetUpdatedAt(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	accessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	baseDeviceBuilder := NewBaseDeviceBuilder(accessTokenBuilder)

	assert.NotNil(t, baseDeviceBuilder)

	baseDeviceBuilder.SetUpdatedAt(1)

	assert.Equal(t, 1, baseDeviceBuilder.device.GetUpdatedAt())

	baseDeviceBuilder.SetUpdatedAt(2)

	assert.Equal(t, 2, baseDeviceBuilder.device.GetUpdatedAt())
}
