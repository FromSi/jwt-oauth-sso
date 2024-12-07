package routes

import (
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLogoutRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)
}

func TestNewLogoutRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)

	assert.Equal(t, "POST", logoutRoute.Method())
}

func TestNewLogoutRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)

	assert.Equal(t, "/logout", logoutRoute.Pattern())
}

func TestNewLogoutRoute_Handle(t *testing.T) {

}
