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

type PasswordResetWithTokenRoute struct {
	resetTokenRepository          repositories.ResetTokenRepository
	resetTokenService             services.ResetTokenService
	userService                   services.UserService
	passwordResetWithTokenRequest requests.PasswordResetWithTokenRequest
	errorBadRequestResponse       responses.ErrorBadRequestResponse
	errorConflictResponse         responses.ErrorConflictResponse
	errorInternalServerResponse   responses.ErrorInternalServerResponse
}

func NewPasswordResetWithTokenRoute(
	resetTokenRepository repositories.ResetTokenRepository,
	resetTokenService services.ResetTokenService,
	userService services.UserService,
	passwordResetWithTokenRequest requests.PasswordResetWithTokenRequest,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorConflictResponse responses.ErrorConflictResponse,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *PasswordResetWithTokenRoute {
	return &PasswordResetWithTokenRoute{
		resetTokenRepository:          resetTokenRepository,
		resetTokenService:             resetTokenService,
		userService:                   userService,
		passwordResetWithTokenRequest: passwordResetWithTokenRequest,
		errorBadRequestResponse:       errorBadRequestResponse,
		errorConflictResponse:         errorConflictResponse,
		errorInternalServerResponse:   errorInternalServerResponse,
	}
}

func (receiver PasswordResetWithTokenRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithTokenRoute) Pattern() string {
	return "/password_reset_with_token"
}

func (receiver PasswordResetWithTokenRoute) Handle(context *gin.Context) {
	request, err := receiver.passwordResetWithTokenRequest.Make(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			receiver.errorBadRequestResponse.Make(err),
		)

		return
	}

	token := receiver.resetTokenRepository.GetActiveResetTokenByToken(request.GetBody().GetToken())

	if token == nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(errors.New("token is expired")),
		)

		return
	}

	hashedPassword, err := receiver.userService.HashPassword(request.GetBody().GetNewPassword())

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	err = receiver.userService.UpdatePasswordByUUIDAndHashedPassword(
		token.GetUserUUID(),
		hashedPassword,
	)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	err = receiver.resetTokenRepository.DeleteResetToken(request.GetBody().GetToken())

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
