package routes

import (
	"errors"
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

func Test_NewLogoutDeviceRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)
}

func TestNewLogoutDeviceRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)

	assert.Equal(t, "POST", logoutDeviceRoute.Method())
}

func TestNewLogoutDeviceRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)

	assert.Equal(t, "/logout_device", logoutDeviceRoute.Pattern())
}

func TestNewLogoutDeviceRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockLogoutDeviceRequest := requests_mocks.NewMockLogoutDeviceRequest(mockController)
	mockLogoutDeviceRequestBody := requests_mocks.NewMockLogoutDeviceRequestBody(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	mockLogoutDeviceRequest.EXPECT().GetBody().Return(mockLogoutDeviceRequestBody).AnyTimes()

	logoutDeviceRoute := NewLogoutDeviceRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutDeviceRequest,
		mockErrorBadRequestResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutDeviceRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(nil, errors.New("error"))

	logoutDeviceRoute.Handle(c)

	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutDeviceRequest.EXPECT().Make(c).Return(nil, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	logoutDeviceRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutDeviceRequest.EXPECT().Make(c).Return(mockLogoutDeviceRequest, nil)
	mockLogoutDeviceRequestBody.EXPECT().GetDeviceUUID().Return("1")
	mockAccessToken.EXPECT().GetSubject().Return("2")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().DeleteDeviceByUUIDAndUserUUID("1", "2").Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	logoutDeviceRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutDeviceRequest.EXPECT().Make(c).Return(mockLogoutDeviceRequest, nil)
	mockLogoutDeviceRequestBody.EXPECT().GetDeviceUUID().Return("1")
	mockAccessToken.EXPECT().GetSubject().Return("2")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().DeleteDeviceByUUIDAndUserUUID("1", "2").Return(nil)

	logoutDeviceRoute.Handle(c)

	assert.Equal(t, http.StatusAccepted, c.Writer.Status())
}
