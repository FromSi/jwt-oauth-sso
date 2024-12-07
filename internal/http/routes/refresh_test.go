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

func Test_NewRefreshRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)
}

func TestNewRefreshRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)

	assert.Equal(t, "POST", refreshRoute.Method())
}

func TestNewRefreshRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)

	assert.Equal(t, "/refresh", refreshRoute.Pattern())
}

func TestNewRefreshRoute_Handle(t *testing.T) {

}
