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

func Test_NewSendResetTokenRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)
}

func TestNewSendResetTokenRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)

	assert.Equal(t, "POST", sendResetTokenRoute.Method())
}

func TestNewSendResetTokenRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)

	assert.Equal(t, "/send_reset_token", sendResetTokenRoute.Pattern())
}

func TestNewSendResetTokenRoute_Handle(t *testing.T) {

}
