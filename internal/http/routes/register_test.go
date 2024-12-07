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

func Test_NewRegisterRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, registerRoute)
}

func TestNewRegisterRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, registerRoute)

	assert.Equal(t, "POST", registerRoute.Method())
}

func TestNewRegisterRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, registerRoute)

	assert.Equal(t, "/register", registerRoute.Pattern())
}

func TestNewRegisterRoute_Handle(t *testing.T) {

}
