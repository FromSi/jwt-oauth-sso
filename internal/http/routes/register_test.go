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

func Test_NewRegisterRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
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
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
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
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, registerRoute)

	assert.Equal(t, "/register", registerRoute.Pattern())
}

func TestNewRegisterRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUser := repositories_mocks.NewMockUser(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockDevice := repositories_mocks.NewMockDevice(mockController)
	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockDeviceService := services_mocks.NewMockDeviceService(mockController)
	mockRegisterRequest := requests_mocks.NewMockRegisterRequest(mockController)
	mockRegisterRequestBody := requests_mocks.NewMockRegisterRequestBody(mockController)
	mockSuccessRegisterResponse := responses_mocks.NewMockSuccessRegisterResponse(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	mockRegisterRequest.EXPECT().GetBody().Return(mockRegisterRequestBody).AnyTimes()

	registerRoute := NewRegisterRoute(
		mockUserRepository,
		mockDeviceRepository,
		mockUserService,
		mockDeviceService,
		mockRegisterRequest,
		mockSuccessRegisterResponse,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, registerRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(nil, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("", errors.New("error"))
	mockErrorConflictResponse.EXPECT().Make(errors.New("error")).Return(mockErrorConflictResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(errors.New("error"))
	mockErrorConflictResponse.EXPECT().Make(errors.New("error")).Return(mockErrorConflictResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(nil)
	mockRegisterRequest.EXPECT().GetIP().Return("2")
	mockRegisterRequest.EXPECT().GetUserAgent().Return("3")
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(nil)
	mockRegisterRequest.EXPECT().GetIP().Return("2")
	mockRegisterRequest.EXPECT().GetUserAgent().Return("3")
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(nil)
	mockRegisterRequest.EXPECT().GetIP().Return("2")
	mockRegisterRequest.EXPECT().GetUserAgent().Return("3")
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(nil)
	mockRegisterRequest.EXPECT().GetIP().Return("2")
	mockRegisterRequest.EXPECT().GetUserAgent().Return("3")
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(nil)
	mockRegisterRequest.EXPECT().GetIP().Return("2")
	mockRegisterRequest.EXPECT().GetUserAgent().Return("3")
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(nil)
	mockSuccessRegisterResponse.EXPECT().Make(mockDevice).Return(nil, errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockRegisterRequest.EXPECT().Make(c).Return(mockRegisterRequest, nil)
	mockRegisterRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)
	mockRegisterRequestBody.EXPECT().GetPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("3", nil)
	mockUserService.EXPECT().GenerateUUID().Return("1")
	mockRegisterRequestBody.EXPECT().GetEmail().Return("2")
	mockUserService.EXPECT().CreateUserByUUIDAndEmailAndHashedPassword("1", "2", "3").Return(nil)
	mockRegisterRequest.EXPECT().GetIP().Return("2")
	mockRegisterRequest.EXPECT().GetUserAgent().Return("3")
	mockDeviceService.EXPECT().GetNewDeviceByUserUUIDAndIpAndUserAgent("1", "2", "3").Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().CreateDevice(mockDevice).Return(nil)
	mockDeviceService.EXPECT().GetNewRefreshDetailsByDevice(mockDevice).Return(mockDevice, nil)
	mockDeviceRepository.EXPECT().UpdateDevice(mockDevice).Return(nil)
	mockSuccessRegisterResponse.EXPECT().Make(mockDevice).Return(mockSuccessRegisterResponse, nil)

	registerRoute.Handle(c)

	assert.Equal(t, http.StatusCreated, c.Writer.Status())
}
