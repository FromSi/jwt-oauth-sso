package routes

import (
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewDevicesRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	devicesRoute := NewDevicesRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, devicesRoute)
}

func TestNewDevicesRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	devicesRoute := NewDevicesRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, devicesRoute)

	assert.Equal(t, "GET", devicesRoute.Method())
}

func TestNewDevicesRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	devicesRoute := NewDevicesRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, devicesRoute)

	assert.Equal(t, "/devices", devicesRoute.Pattern())
}

func TestNewDevicesRoute_Handle(t *testing.T) {

}
