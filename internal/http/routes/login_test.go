package routes

import (
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewLoginRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
	)

	assert.NotEmpty(t, loginRoute)
}

func TestNewLoginRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
	)

	assert.NotEmpty(t, loginRoute)

	assert.Equal(t, "POST", loginRoute.Method())
}

func TestNewLoginRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
	)

	assert.NotEmpty(t, loginRoute)

	assert.Equal(t, "/login", loginRoute.Pattern())
}

func TestNewLoginRoute_Handle(t *testing.T) {

}
