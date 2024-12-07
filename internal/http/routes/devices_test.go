package routes

import (
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewDevicesRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)
}

func TestNewDevicesRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)

	assert.Equal(t, "GET", devicesRoute.Method())
}

func TestNewDevicesRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)

	assert.Equal(t, "/devices", devicesRoute.Pattern())
}

func TestNewDevicesRoute_Handle(t *testing.T) {

}
