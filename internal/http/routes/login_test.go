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

func Test_NewLoginRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)
}

func TestNewLoginRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)

	assert.Equal(t, "POST", loginRoute.Method())
}

func TestNewLoginRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)

	assert.Equal(t, "/login", loginRoute.Pattern())
}

func TestNewLoginRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockUser := repositories_mocks.NewMockUser(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDevice := repositories_mocks.NewMockDevice(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockLoginRequest := requests_mocks.NewMockLoginRequest(mockController)
	mockLoginRequestBody := requests_mocks.NewMockLoginRequestBody(mockController)
	mockSuccessLoginResponse := responses_mocks.NewMockSuccessLoginResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	mockLoginRequest.EXPECT().GetBody().Return(mockLoginRequestBody).AnyTimes()

	loginRoute := NewLoginRoute(
		mockUserService,
		mockDeviceService,
		mockUserRepository,
		mockDeviceRepository,
		mockLoginRequest,
		mockSuccessLoginResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, loginRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(nil, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(errors.New("error"))
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()
	mockLoginRequest.EXPECT().GetIP().Return("2").AnyTimes()
	mockLoginRequest.EXPECT().GetUserAgent().Return("3").AnyTimes()
	mockDeviceService.EXPECT().GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil)
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()
	mockLoginRequest.EXPECT().GetIP().Return("2").AnyTimes()
	mockLoginRequest.EXPECT().GetUserAgent().Return("3").AnyTimes()
	mockDeviceService.EXPECT().GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil)
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()
	mockLoginRequest.EXPECT().GetIP().Return("2").AnyTimes()
	mockLoginRequest.EXPECT().GetUserAgent().Return("3").AnyTimes()
	mockDeviceService.EXPECT().GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil)
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()
	mockLoginRequest.EXPECT().GetIP().Return("2").AnyTimes()
	mockLoginRequest.EXPECT().GetUserAgent().Return("3").AnyTimes()
	mockDeviceService.EXPECT().GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil)
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()
	mockLoginRequest.EXPECT().GetIP().Return("2").AnyTimes()
	mockLoginRequest.EXPECT().GetUserAgent().Return("3").AnyTimes()
	mockDeviceService.EXPECT().GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil)
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(nil)
	mockSuccessLoginResponse.EXPECT().Make(mockDevice).Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockLoginRequest.EXPECT().Make(c).Return(mockLoginRequest, nil)
	mockLoginRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockUser.EXPECT().GetPassword().Return("1")
	mockLoginRequestBody.EXPECT().GetPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUser.EXPECT().GetUUID().Return("1").AnyTimes()
	mockLoginRequest.EXPECT().GetIP().Return("2").AnyTimes()
	mockLoginRequest.EXPECT().GetUserAgent().Return("3").AnyTimes()
	mockDeviceService.EXPECT().GetOldDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil)
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(nil)
	mockSuccessLoginResponse.EXPECT().Make(mockDevice).Return(mockSuccessLoginResponse, nil)

	loginRoute.Handle(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
