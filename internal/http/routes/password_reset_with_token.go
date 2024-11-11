package routes

import (
	"errors"
	"fmt"
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
}

func NewPasswordResetWithTokenRoute(
	resetTokenRepository repositories.ResetTokenRepository,
	resetTokenService services.ResetTokenService,
) *PasswordResetWithTokenRoute {
	return &PasswordResetWithTokenRoute{
		resetTokenRepository: resetTokenRepository,
		resetTokenService:    resetTokenService,
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

	fmt.Println(map[string]any{
		"token":        request.Body.Token,
		"new_password": request.Body.NewPassword,
	})

	taken := receiver.resetTokenRepository.GetActiveResetTokenByToken(request.Body.Token)

	if taken == nil {
		context.JSON(
			http.StatusConflict,
			responses.NewErrorConflictResponse(errors.New("token is expired")),
		)

		return
	}

	err := receiver.resetTokenService.ResetPasswordByUserUUIDAndNewPassword(
		taken.GetUserUUID(),
		request.Body.NewPassword,
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
