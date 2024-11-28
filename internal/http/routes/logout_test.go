package routes

import (
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLogoutRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutRoute := NewLogoutRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutRoute)
}

func TestNewLogoutRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutRoute := NewLogoutRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutRoute)

	assert.Equal(t, "POST", logoutRoute.Method())
}

func TestNewLogoutRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	logoutRoute := NewLogoutRoute(mockDeviceRepository, mockAccessTokenBuilder)

	assert.NotEmpty(t, logoutRoute)

	assert.Equal(t, "/logout", logoutRoute.Pattern())
}

func TestNewLogoutRoute_Handle(t *testing.T) {

}
