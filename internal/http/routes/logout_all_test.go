package routes

import (
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLogoutAllRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockAccessTokenBuilder,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)
}

func TestNewLogoutAllRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockAccessTokenBuilder,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)

	assert.Equal(t, "POST", logoutAllRoute.Method())
}

func TestNewLogoutAllRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockAccessTokenBuilder,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)

	assert.Equal(t, "/logout_all", logoutAllRoute.Pattern())
}

func TestNewLogoutAllRoute_Handle(t *testing.T) {

}
