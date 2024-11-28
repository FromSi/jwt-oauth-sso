package routes

import (
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLogoutDeviceRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutDeviceRoute)
}

func TestNewLogoutDeviceRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutDeviceRoute)

	assert.Equal(t, "POST", logoutDeviceRoute.Method())
}

func TestNewLogoutDeviceRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutDeviceRoute)

	assert.Equal(t, "/logout_device", logoutDeviceRoute.Pattern())
}

func TestNewLogoutDeviceRoute_Handle(t *testing.T) {

}
