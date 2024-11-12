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
	resetTokenRepository repositories.ResetTokenRepository
	resetTokenService    services.ResetTokenService
	userService          services.UserService
}

func NewPasswordResetWithTokenRoute(
	resetTokenRepository repositories.ResetTokenRepository,
	resetTokenService services.ResetTokenService,
	userService services.UserService,
) *PasswordResetWithTokenRoute {
	return &PasswordResetWithTokenRoute{
		resetTokenRepository: resetTokenRepository,
		resetTokenService:    resetTokenService,
		userService:          userService,
	}
}

func (receiver PasswordResetWithTokenRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithTokenRoute) Pattern() string {
	return "/password_reset_with_token"
}

func (receiver PasswordResetWithTokenRoute) Handle(context *gin.Context) {
	request, errResponse := requests.NewPasswordResetWithTokenRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	token := receiver.resetTokenRepository.GetActiveResetTokenByToken(request.Body.Token)

	if token == nil {
		context.JSON(
			http.StatusConflict,
			responses.NewErrorConflictResponse(errors.New("token is expired")),
		)

		return
	}

	hashedPassword, err := receiver.userService.HashPassword(request.Body.NewPassword)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
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
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	err = receiver.resetTokenRepository.DeleteResetToken(request.Body.Token)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
