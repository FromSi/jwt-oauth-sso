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

func Test_NewPasswordResetWithTokenRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)
}

func TestNewPasswordResetWithTokenRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)

	assert.Equal(t, "POST", passwordResetWithTokenRoute.Method())
}

func TestNewPasswordResetWithTokenRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)

	assert.Equal(t, "/password_reset_with_token", passwordResetWithTokenRoute.Pattern())
}

func TestNewPasswordResetWithTokenRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockResetToken := repositories_mocks.NewMockResetToken(mockController)
	mockResetTokenRepository := repositories_mocks.NewMockResetTokenRepository(mockController)
	mockUserService := services_mocks.NewMockUserService(mockController)
	mockPasswordResetWithTokenRequest := requests_mocks.NewMockPasswordResetWithTokenRequest(mockController)
	mockPasswordResetWithTokenRequestBody := requests_mocks.NewMockPasswordResetWithTokenRequestBody(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)
	mockErrorInternalServerResponse := responses_mocks.NewMockErrorInternalServerResponse(mockController)

	mockPasswordResetWithTokenRequest.EXPECT().GetBody().Return(mockPasswordResetWithTokenRequestBody).AnyTimes()

	passwordResetWithTokenRoute := NewPasswordResetWithTokenRoute(
		mockResetTokenRepository,
		mockUserService,
		mockPasswordResetWithTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
		mockErrorInternalServerResponse,
	)

	assert.NotEmpty(t, passwordResetWithTokenRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockPasswordResetWithTokenRequest.EXPECT().Make(c).Return(nil, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	passwordResetWithTokenRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockPasswordResetWithTokenRequest.EXPECT().Make(c).Return(mockPasswordResetWithTokenRequest, nil)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockResetTokenRepository.EXPECT().GetActiveResetTokenByToken("1").Return(nil)
	mockErrorConflictResponse.EXPECT().Make(gomock.Any()).Return(mockErrorConflictResponse)

	passwordResetWithTokenRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockPasswordResetWithTokenRequest.EXPECT().Make(c).Return(mockPasswordResetWithTokenRequest, nil)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockResetTokenRepository.EXPECT().GetActiveResetTokenByToken("1").Return(mockResetToken)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetNewPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("", errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	passwordResetWithTokenRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockPasswordResetWithTokenRequest.EXPECT().Make(c).Return(mockPasswordResetWithTokenRequest, nil)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockResetToken.EXPECT().GetUserUUID().Return("1")
	mockResetTokenRepository.EXPECT().GetActiveResetTokenByToken("1").Return(mockResetToken)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetNewPassword().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("2", nil)
	mockUserService.EXPECT().UpdatePasswordByUUIDAndHashedPassword("1", "2").Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	passwordResetWithTokenRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockPasswordResetWithTokenRequest.EXPECT().Make(c).Return(mockPasswordResetWithTokenRequest, nil)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockResetToken.EXPECT().GetUserUUID().Return("1")
	mockResetTokenRepository.EXPECT().GetActiveResetTokenByToken("1").Return(mockResetToken)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetNewPassword().Return("1")
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("2", nil)
	mockUserService.EXPECT().UpdatePasswordByUUIDAndHashedPassword("1", "2").Return(nil)
	mockResetTokenRepository.EXPECT().DeleteResetToken("1").Return(errors.New("error"))
	mockErrorInternalServerResponse.EXPECT().Make(errors.New("error")).Return(mockErrorInternalServerResponse)

	passwordResetWithTokenRoute.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockPasswordResetWithTokenRequest.EXPECT().Make(c).Return(mockPasswordResetWithTokenRequest, nil)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockResetToken.EXPECT().GetUserUUID().Return("1")
	mockResetTokenRepository.EXPECT().GetActiveResetTokenByToken("1").Return(mockResetToken)
	mockPasswordResetWithTokenRequestBody.EXPECT().GetNewPassword().Return("1")
	mockPasswordResetWithTokenRequestBody.EXPECT().GetToken().Return("1")
	mockUserService.EXPECT().HashPassword("1").Return("2", nil)
	mockUserService.EXPECT().UpdatePasswordByUUIDAndHashedPassword("1", "2").Return(nil)
	mockResetTokenRepository.EXPECT().DeleteResetToken("1").Return(nil)

	passwordResetWithTokenRoute.Handle(c)

	assert.Equal(t, http.StatusAccepted, c.Writer.Status())
}
