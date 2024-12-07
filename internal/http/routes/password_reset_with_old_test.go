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

func Test_NewPasswordResetWithOldRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)
}

func TestNewPasswordResetWithOldRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	assert.Equal(t, "POST", passwordResetWithOldRoute.Method())
}

func TestNewPasswordResetWithOldRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	assert.Equal(t, "/password_reset_with_old", passwordResetWithOldRoute.Pattern())
}

func TestNewPasswordResetWithOldRoute_Handle(t *testing.T) {

}
