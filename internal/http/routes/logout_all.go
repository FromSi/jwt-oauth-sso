package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutAllRoute struct {
	deviceRepository            repositories.DeviceRepository
	accessTokenBuilder          tokens.AccessTokenBuilder
	bearerAuthRequestHeader     requests.BearerAuthRequestHeader
	logoutAllRequest            requests.LogoutAllRequest
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewLogoutAllRoute(
	deviceRepository repositories.DeviceRepository,
	accessTokenBuilder tokens.AccessTokenBuilder,
	bearerAuthRequestHeader requests.BearerAuthRequestHeader,
	logoutAllRequest requests.LogoutAllRequest,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *LogoutAllRoute {
	return &LogoutAllRoute{
		deviceRepository:            deviceRepository,
		accessTokenBuilder:          accessTokenBuilder,
		bearerAuthRequestHeader:     bearerAuthRequestHeader,
		logoutAllRequest:            logoutAllRequest,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver LogoutAllRoute) Method() string {
	return "POST"
}

func (receiver LogoutAllRoute) Pattern() string {
	return "/logout_all"
}

func (receiver LogoutAllRoute) Handle(context *gin.Context) {
	headers, err := receiver.bearerAuthRequestHeader.Make(context)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = receiver.logoutAllRequest.Make(context)

	err = receiver.deviceRepository.DeleteAllDevicesByUserUUID(headers.GetAccessToken().GetSubject())

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
