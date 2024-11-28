package routes

import (
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewPasswordResetWithOldRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockAccessTokenBuilder,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)
}

func TestNewPasswordResetWithOldRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockAccessTokenBuilder,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	assert.Equal(t, "POST", passwordResetWithOldRoute.Method())
}

func TestNewPasswordResetWithOldRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockAccessTokenBuilder,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	assert.Equal(t, "/password_reset_with_old", passwordResetWithOldRoute.Pattern())
}

func TestNewPasswordResetWithOldRoute_Handle(t *testing.T) {

}
