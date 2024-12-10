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

func Test_NewLogoutAllRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)
}

func TestNewLogoutAllRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)

	assert.Equal(t, "POST", logoutAllRoute.Method())
}

func TestNewLogoutAllRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)

	assert.Equal(t, "/logout_all", logoutAllRoute.Pattern())
}

func TestNewLogoutAllRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDeviceRepository := repositories_mocks.NewMockDeviceRepository(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockLogoutAllRequest := requests_mocks.NewMockLogoutAllRequest(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	logoutAllRoute := NewLogoutAllRoute(
		mockDeviceRepository,
		mockBearerAuthRequestHeader,
		mockLogoutAllRequest,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, logoutAllRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(nil, errors.New("error"))

	logoutAllRoute.Handle(c)

	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutAllRequest.EXPECT().Make(c).Return(nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().DeleteAllDevicesByUserUUID("1").Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	logoutAllRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockLogoutAllRequest.EXPECT().Make(c).Return(nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockDeviceRepository.EXPECT().DeleteAllDevicesByUserUUID("1").Return(nil)

	logoutAllRoute.Handle(c)

	assert.Equal(t, http.StatusAccepted, c.Writer.Status())
}
