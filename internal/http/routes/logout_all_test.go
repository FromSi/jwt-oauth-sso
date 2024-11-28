package routes

import (
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

	logoutAllRoute := NewLogoutAllRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutAllRoute)
}

func TestNewLogoutAllRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutAllRoute := NewLogoutAllRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutAllRoute)

	assert.Equal(t, "POST", logoutAllRoute.Method())
}

func TestNewLogoutAllRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutAllRoute := NewLogoutAllRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutAllRoute)

	assert.Equal(t, "/logout_all", logoutAllRoute.Pattern())
}

func TestNewLogoutAllRoute_Handle(t *testing.T) {

}
