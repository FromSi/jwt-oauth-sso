package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutRoute struct {
	deviceRepository            repositories.DeviceRepository
	bearerAuthRequestHeader     requests.BearerAuthRequestHeader
	logoutRequest               requests.LogoutRequest
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewLogoutRoute(
	deviceRepository repositories.DeviceRepository,
	bearerAuthRequestHeader requests.BearerAuthRequestHeader,
	logoutRequest requests.LogoutRequest,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *LogoutRoute {
	return &LogoutRoute{
		deviceRepository:            deviceRepository,
		bearerAuthRequestHeader:     bearerAuthRequestHeader,
		logoutRequest:               logoutRequest,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver LogoutRoute) Method() string {
	return "POST"
}

func (receiver LogoutRoute) Pattern() string {
	return "/logout"
}

func (receiver LogoutRoute) Handle(context *gin.Context) {
	headers, err := receiver.bearerAuthRequestHeader.Make(context)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = receiver.logoutRequest.Make(context)

	err = receiver.deviceRepository.DeleteDeviceByUUID(headers.GetAccessToken().GetDeviceUUID())

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
