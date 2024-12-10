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

func Test_NewSendResetTokenRoute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)
}

func TestNewSendResetTokenRoute_Method(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)

	assert.Equal(t, "POST", sendResetTokenRoute.Method())
}

func TestNewSendResetTokenRoute_Pattern(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)

	assert.Equal(t, "/send_reset_token", sendResetTokenRoute.Pattern())
}

func TestNewSendResetTokenRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUser := repositories_mocks.NewMockUser(mockController)
	mockUserRepository := repositories_mocks.NewMockUserRepository(mockController)
	mockResetTokenService := services_mocks.NewMockResetTokenService(mockController)
	mockSendResetTokenRequest := requests_mocks.NewMockSendResetTokenRequest(mockController)
	mockSendResetTokenRequestBody := requests_mocks.NewMockSendResetTokenRequestBody(mockController)
	mockErrorBadRequestResponse := responses_mocks.NewMockErrorBadRequestResponse(mockController)
	mockErrorConflictResponse := responses_mocks.NewMockErrorConflictResponse(mockController)

	mockSendResetTokenRequest.EXPECT().GetBody().Return(mockSendResetTokenRequestBody).AnyTimes()

	sendResetTokenRoute := NewSendResetTokenRoute(
		mockUserRepository,
		mockResetTokenService,
		mockSendResetTokenRequest,
		mockErrorBadRequestResponse,
		mockErrorConflictResponse,
	)

	assert.NotEmpty(t, sendResetTokenRoute)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockSendResetTokenRequest.EXPECT().Make(c).Return(nil, errors.New("error"))
	mockErrorBadRequestResponse.EXPECT().Make(errors.New("error")).Return(mockErrorBadRequestResponse)

	sendResetTokenRoute.Handle(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockSendResetTokenRequest.EXPECT().Make(c).Return(mockSendResetTokenRequest, nil)
	mockSendResetTokenRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(nil)

	sendResetTokenRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockSendResetTokenRequest.EXPECT().Make(c).Return(mockSendResetTokenRequest, nil)
	mockSendResetTokenRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockResetTokenService.EXPECT().SendNewResetTokenByUser(mockUser).Return(errors.New("error"))
	mockErrorConflictResponse.EXPECT().Make(errors.New("error")).Return(mockErrorConflictResponse)

	sendResetTokenRoute.Handle(c)

	assert.Equal(t, http.StatusConflict, c.Writer.Status())

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	mockSendResetTokenRequest.EXPECT().Make(c).Return(mockSendResetTokenRequest, nil)
	mockSendResetTokenRequestBody.EXPECT().GetEmail().Return("1")
	mockUserRepository.EXPECT().GetUserByEmail("1").Return(mockUser)
	mockResetTokenService.EXPECT().SendNewResetTokenByUser(mockUser).Return(nil)

	sendResetTokenRoute.Handle(c)

	assert.Equal(t, http.StatusAccepted, c.Writer.Status())
}
