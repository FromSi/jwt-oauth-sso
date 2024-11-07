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
	userRepository    repositories.UserRepository
	resetTokenService services.ResetTokenService
}

func NewSendResetTokenRoute(
	userRepository repositories.UserRepository,
	resetTokenService services.ResetTokenService,
) *SendResetTokenRoute {
	return &SendResetTokenRoute{
		userRepository:    userRepository,
		resetTokenService: resetTokenService,
	}
}

func (receiver SendResetTokenRoute) Method() string {
	return "POST"
}

func (receiver SendResetTokenRoute) Pattern() string {
	return "/send_reset_token"
}

func (receiver SendResetTokenRoute) Handle(context *gin.Context) {
	request, errResponse := requests.NewSendResetTokenRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	userExists := receiver.userRepository.HasUserByEmail(request.Body.Email)

	if !userExists {
		err := responses.NewErrorConflictResponse(
			errors.New("user not found with this email"),
		)

		context.JSON(http.StatusConflict, err)

		return
	}

	err := receiver.resetTokenService.SendNewResetTokenByUserEmail(request.Body.Email)

	if err != nil {
		context.JSON(http.StatusConflict, responses.NewErrorConflictResponse(err))

		return
	}

	context.Status(http.StatusAccepted)
}
