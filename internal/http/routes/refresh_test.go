package routes

import (
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_NewRefreshRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)

	refreshRoute := NewRefreshRoute(mockDeviceService, mockDeviceRepository)

	assert.NotEmpty(t, refreshRoute)
}

func TestNewRefreshRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)

	refreshRoute := NewRefreshRoute(mockDeviceService, mockDeviceRepository)

	assert.NotEmpty(t, refreshRoute)

	assert.Equal(t, "POST", refreshRoute.Method())
}

func TestNewRefreshRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)

	refreshRoute := NewRefreshRoute(mockDeviceService, mockDeviceRepository)

	assert.NotEmpty(t, refreshRoute)

	assert.Equal(t, "/refresh", refreshRoute.Pattern())
}

func TestNewRefreshRoute_Handle(t *testing.T) {

}
