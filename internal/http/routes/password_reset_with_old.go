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

type PasswordResetWithOldRoute struct {
	userRepository              repositories.UserRepository
	userService                 services.UserService
	bearerAuthRequestHeader     requests.BearerAuthRequestHeader
	passwordResetWithOldRequest requests.PasswordResetWithOldRequest
	errorBadRequestResponse     responses.ErrorBadRequestResponse
	errorConflictResponse       responses.ErrorConflictResponse
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewPasswordResetWithOldRoute(
	userRepository repositories.UserRepository,
	userService services.UserService,
	bearerAuthRequestHeader requests.BearerAuthRequestHeader,
	passwordResetWithOldRequest requests.PasswordResetWithOldRequest,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorConflictResponse responses.ErrorConflictResponse,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *PasswordResetWithOldRoute {
	return &PasswordResetWithOldRoute{
		userRepository:              userRepository,
		userService:                 userService,
		bearerAuthRequestHeader:     bearerAuthRequestHeader,
		passwordResetWithOldRequest: passwordResetWithOldRequest,
		errorBadRequestResponse:     errorBadRequestResponse,
		errorConflictResponse:       errorConflictResponse,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver PasswordResetWithOldRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithOldRoute) Pattern() string {
	return "/password_reset_with_old"
}

func (receiver PasswordResetWithOldRoute) Handle(context *gin.Context) {
	headers, err := receiver.bearerAuthRequestHeader.Make(context)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, err := receiver.passwordResetWithOldRequest.Make(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			receiver.errorBadRequestResponse.Make(err),
		)

		return
	}

	user := receiver.userRepository.GetUserByUUID(headers.GetAccessToken().GetSubject())

	if user == nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(errors.New("user not found")),
		)

		return
	}

	err = receiver.userService.CheckHashedPasswordAndNativePassword(
		user.GetPassword(),
		request.GetBody().GetOldPassword(),
	)

	if err != nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(errors.New("invalid old password")),
		)

		return
	}

	err = receiver.userService.UpdatePasswordByUUIDAndHashedPassword(
		user.GetUUID(),
		request.GetBody().GetNewPassword(),
	)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
