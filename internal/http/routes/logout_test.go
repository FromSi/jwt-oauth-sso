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

func Test_NewLogoutRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)
}

func TestNewLogoutRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)

	assert.Equal(t, "POST", logoutRoute.Method())
}

func TestNewLogoutRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)

	assert.Equal(t, "/logout", logoutRoute.Pattern())
}

func TestNewLogoutRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockLogoutRequest := requests_mocks.NewMockLogoutRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutRoute := NewLogoutRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(nil, errors.New("error"))

	logoutRoute.Handle(c)

	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutRequest.EXPECT().Make(c).Return(nil)
	mockAccessToken.EXPECT().GetDeviceUUID().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().DeleteDeviceByUUID("1").Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	logoutRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutRequest.EXPECT().Make(c).Return(nil)
	mockAccessToken.EXPECT().GetDeviceUUID().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().DeleteDeviceByUUID("1").Return(nil)

	logoutRoute.Handle(c)

	assert.Equal(t, http.StatusAccepted, c.Writer.Status())
}
