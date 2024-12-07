package routes

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendResetTokenRoute struct {
	userRepository          repositories.UserRepository
	resetTokenService       services.ResetTokenService
	sendResetTokenRequest   requests.SendResetTokenRequest
	errorBadRequestResponse responses.ErrorBadRequestResponse
	errorConflictResponse   responses.ErrorConflictResponse
}

func NewSendResetTokenRoute(
	userRepository repositories.UserRepository,
	resetTokenService services.ResetTokenService,
	sendResetTokenRequest requests.SendResetTokenRequest,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorConflictResponse responses.ErrorConflictResponse,
) *SendResetTokenRoute {
	return &SendResetTokenRoute{
		userRepository:          userRepository,
		resetTokenService:       resetTokenService,
		sendResetTokenRequest:   sendResetTokenRequest,
		errorBadRequestResponse: errorBadRequestResponse,
		errorConflictResponse:   errorConflictResponse,
	}
}

func (receiver SendResetTokenRoute) Method() string {
	return "POST"
}

func (receiver SendResetTokenRoute) Pattern() string {
	return "/send_reset_token"
}

func (receiver SendResetTokenRoute) Handle(context *gin.Context) {
	request, err := receiver.sendResetTokenRequest.Make(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			receiver.errorBadRequestResponse.Make(err),
		)

		return
	}

	user := receiver.userRepository.GetUserByEmail(request.GetBody().GetEmail())

	if user == nil {
		context.JSON(
			http.StatusConflict,
			errors.New("user not found with this email"),
		)

		return
	}

	err = receiver.resetTokenService.SendNewResetTokenByUser(user)

	if err != nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
