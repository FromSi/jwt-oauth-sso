package routes

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewDevicesRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)
}

func TestNewDevicesRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)

	assert.Equal(t, "GET", devicesRoute.Method())
}

func TestNewDevicesRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)

	assert.Equal(t, "/devices", devicesRoute.Pattern())
}

func TestNewDevicesRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockDevicesRequest := requests_mocks.NewMockDevicesRequest(mockController)
	mockSuccessDevicesResponse := responses_mocks.NewMockSuccessDevicesResponse(mockController)

	devicesRoute := NewDevicesRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockDevicesRequest,
		mockSuccessDevicesResponse,
	)

	assert.NotEmpty(t, devicesRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(nil, errors.New("error"))

	devicesRoute.Handle(c)

	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockDevicesRequest.EXPECT().Make(c).Return(nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().GetDevicesByUserUUID("1").Return([]repositories.Device{})
	mockSuccessDevicesResponse.EXPECT().Make([]repositories.Device{}).Return(mockSuccessDevicesResponse)

	devicesRoute.Handle(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
