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

func Test_NewPasswordResetWithTokenRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockResetTokenService,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)
}

func TestNewPasswordResetWithTokenRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockResetTokenService,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)

	assert.Equal(t, "POST", passwordResetWithTokenRoute.Method())
}

func TestNewPasswordResetWithTokenRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockResetTokenService,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)

	assert.Equal(t, "/password_reset_with_token", passwordResetWithTokenRoute.Pattern())
}

func TestNewPasswordResetWithTokenRoute_Handle(t *testing.T) {

}
