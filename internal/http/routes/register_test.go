package routes

import (
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

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
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

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
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

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
	)

	assert.NotEmpty(t, registerRoute)

	assert.Equal(t, "/register", registerRoute.Pattern())
}

func TestNewRegisterRoute_Handle(t *testing.T) {

}
