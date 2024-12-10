package routes

import (
	"errors"
	requests_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/requests"
	responses_mocks "github.com/fromsi/jwt-oauth-sso/mocks/http/responses"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	services_mocks "github.com/fromsi/jwt-oauth-sso/mocks/services"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewPasswordResetWithOldRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)
}

func TestNewPasswordResetWithOldRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	assert.Equal(t, "POST", passwordResetWithOldRoute.Method())
}

func TestNewPasswordResetWithOldRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	assert.Equal(t, "/password_reset_with_old", passwordResetWithOldRoute.Pattern())
}

func TestNewPasswordResetWithOldRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUser := repositories_mocks.NewMockUser(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockBearerAuthRequestHeader := requests_mocks.NewMockBearerAuthRequestHeader(mockController)
	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockPasswordResetWithOldRequest := requests_mocks.NewMockPasswordResetWithOldRequest(mockController)
	mockPasswordResetWithOldRequestBody := requests_mocks.NewMockPasswordResetWithOldRequestBody(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	mockPasswordResetWithOldRequest.EXPECT().GetBody().Return(mockPasswordResetWithOldRequestBody).AnyTimes()

	passwordResetWithOldRoute := NewPasswordResetWithOldRoute(
		mockUserRepository,
		mockUserService,
		mockBearerAuthRequestHeader,
		mockPasswordResetWithOldRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithOldRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(nil, errors.New("error"))

	passwordResetWithOldRoute.Handle(c)

	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockPasswordResetWithOldRequest.EXPECT().Make(c).Return(nil, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	passwordResetWithOldRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockPasswordResetWithOldRequest.EXPECT().Make(c).Return(mockPasswordResetWithOldRequest, nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockUserRepository.EXPECT().GetUserByUUID("1").Return(nil)
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	passwordResetWithOldRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockPasswordResetWithOldRequest.EXPECT().Make(c).Return(mockPasswordResetWithOldRequest, nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockUser.EXPECT().GetPassword().Return("1")
	mockUserRepository.EXPECT().GetUserByUUID("1").Return(mockUser)
	mockPasswordResetWithOldRequestBody.EXPECT().GetOldPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(errors.New("error"))
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	passwordResetWithOldRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockPasswordResetWithOldRequest.EXPECT().Make(c).Return(mockPasswordResetWithOldRequest, nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockUser.EXPECT().GetPassword().Return("1")
	mockUser.EXPECT().GetUUID().Return("1")
	mockUserRepository.EXPECT().GetUserByUUID("1").Return(mockUser)
	mockPasswordResetWithOldRequestBody.EXPECT().GetOldPassword().Return("2")
	mockPasswordResetWithOldRequestBody.EXPECT().GetNewPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUserService.EXPECT().UpdatePasswordByUUIDAndHashedPassword("1", "2").Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	passwordResetWithOldRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockBearerAuthRequestHeader.EXPECT().Make(c).Return(mockBearerAuthRequestHeader, nil)
	mockPasswordResetWithOldRequest.EXPECT().Make(c).Return(mockPasswordResetWithOldRequest, nil)
	mockAccessToken.EXPECT().GetSubject().Return("1")
	mockBearerAuthRequestHeader.EXPECT().GetAccessToken().Return(mockAccessToken)
	mockUser.EXPECT().GetPassword().Return("1")
	mockUser.EXPECT().GetUUID().Return("1")
	mockUserRepository.EXPECT().GetUserByUUID("1").Return(mockUser)
	mockPasswordResetWithOldRequestBody.EXPECT().GetOldPassword().Return("2")
	mockPasswordResetWithOldRequestBody.EXPECT().GetNewPassword().Return("2")
	mockUserService.EXPECT().CheckHashedPasswordAndNativePassword("1", "2").Return(nil)
	mockUserService.EXPECT().UpdatePasswordByUUIDAndHashedPassword("1", "2").Return(nil)

	passwordResetWithOldRoute.Handle(c)

	assert.Equal(t, http.StatusAccepted, c.Writer.Status())
}
