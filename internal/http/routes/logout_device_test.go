package routes

import (
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLogoutDeviceRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)
}

func TestNewLogoutDeviceRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)

	assert.Equal(t, "POST", logoutDeviceRoute.Method())
}

func TestNewLogoutDeviceRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)

	assert.Equal(t, "/logout_device", logoutDeviceRoute.Pattern())
}

func TestNewLogoutDeviceRoute_Handle(t *testing.T) {

}
