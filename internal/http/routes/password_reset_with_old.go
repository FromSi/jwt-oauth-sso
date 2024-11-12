package routes

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PasswordResetWithOldRoute struct {
	config         configs.TokenConfig
	userRepository repositories.UserRepository
	userService    services.UserService
}

func NewPasswordResetWithOldRoute(
	config configs.TokenConfig,
	userRepository repositories.UserRepository,
	userService services.UserService,
) *PasswordResetWithOldRoute {
	return &PasswordResetWithOldRoute{
		config:         config,
		userRepository: userRepository,
		userService:    userService,
	}
}

func (receiver PasswordResetWithOldRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithOldRoute) Pattern() string {
	return "/password_reset_with_old"
}

func (receiver PasswordResetWithOldRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, errResponse := requests.NewPasswordResetWithOldRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	user := receiver.userRepository.GetUserByUUID(headers.AccessToken.Subject)

	if user == nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(errors.New("user not found")),
		)

		return
	}

	err = receiver.userService.CheckHashedPasswordAndNativePassword(
		user.GetPassword(),
		request.Body.OldPassword,
	)

	if err != nil {
		context.JSON(
			http.StatusConflict,
			responses.NewErrorConflictResponse(errors.New("invalid old password")),
		)

		return
	}

	err = receiver.userService.UpdatePasswordByUUIDAndHashedPassword(
		user.GetUUID(),
		request.Body.NewPassword,
	)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
