package routes

import (
	"errors"
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewRefreshRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)
}

func TestNewRefreshRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)

	assert.Equal(t, "POST", refreshRoute.Method())
}

func TestNewRefreshRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)

	assert.Equal(t, "/refresh", refreshRoute.Pattern())
}

func TestNewRefreshRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockDevice := repositories_mocks.NewMockDevice(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockRefreshRequest := requests_mocks.NewMockRefreshRequest(mockController)
	mockRefreshRequestBody := requests_mocks.NewMockRefreshRequestBody(mockController)
	mockSuccessRefreshResponse := responses_mocks.NewMockSuccessRefreshResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	mockRefreshRequest.EXPECT().GetBody().Return(mockRefreshRequestBody).AnyTimes()

	refreshRoute := NewRefreshRoute(
		mockDeviceService,
		mockDeviceRepository,
		mockRefreshRequest,
		mockSuccessRefreshResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, refreshRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockRefreshRequest.EXPECT().Make(c).Return(mockRefreshRequest, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	refreshRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRefreshRequest.EXPECT().Make(c).Return(mockRefreshRequest, nil)
	mockRefreshRequestBody.EXPECT().GetRefreshToken().Return("1")
	mockDeviceRepository.EXPECT().GetDeviceByRefreshToken("1").Return(nil)
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	refreshRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRefreshRequest.EXPECT().Make(c).Return(mockRefreshRequest, nil)
	mockRefreshRequestBody.EXPECT().GetRefreshToken().Return("1")
	mockDeviceRepository.EXPECT().GetDeviceByRefreshToken("1").Return(mockDevice)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	refreshRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRefreshRequest.EXPECT().Make(c).Return(mockRefreshRequest, nil)
	mockRefreshRequestBody.EXPECT().GetRefreshToken().Return("1")
	mockDeviceRepository.EXPECT().GetDeviceByRefreshToken("1").Return(mockDevice)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	refreshRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRefreshRequest.EXPECT().Make(c).Return(mockRefreshRequest, nil)
	mockRefreshRequestBody.EXPECT().GetRefreshToken().Return("1")
	mockDeviceRepository.EXPECT().GetDeviceByRefreshToken("1").Return(mockDevice)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(nil)
	mockSuccessRefreshResponse.EXPECT().Make(mockDevice).Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	refreshRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRefreshRequest.EXPECT().Make(c).Return(mockRefreshRequest, nil)
	mockRefreshRequestBody.EXPECT().GetRefreshToken().Return("1")
	mockDeviceRepository.EXPECT().GetDeviceByRefreshToken("1").Return(mockDevice)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(nil)
	mockSuccessRefreshResponse.EXPECT().Make(mockDevice).Return(mockSuccessRefreshResponse, nil)

	refreshRoute.Handle(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
