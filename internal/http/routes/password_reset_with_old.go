package routes

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PasswordResetWithOldRoute struct {
	userRepository     repositories.UserRepository
	userService        services.UserService
	accessTokenBuilder tokens.AccessTokenBuilder
}

func NewPasswordResetWithOldRoute(
	userRepository repositories.UserRepository,
	userService services.UserService,
	accessTokenBuilder tokens.AccessTokenBuilder,
) *PasswordResetWithOldRoute {
	return &PasswordResetWithOldRoute{
		userRepository:     userRepository,
		userService:        userService,
		accessTokenBuilder: accessTokenBuilder,
	}
}

func (receiver PasswordResetWithOldRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithOldRoute) Pattern() string {
	return "/password_reset_with_old"
}

func (receiver PasswordResetWithOldRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.accessTokenBuilder)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, errResponse := requests.NewPasswordResetWithOldRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	user := receiver.userRepository.GetUserByUUID(headers.AccessToken.GetSubject())

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
