package routes

import (
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLoginRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)
}

func TestNewLoginRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)

	assert.Equal(t, "POST", loginRoute.Method())
}

func TestNewLoginRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)

	assert.Equal(t, "/login", loginRoute.Pattern())
}

func TestNewLoginRoute_Handle(t *testing.T) {

}
