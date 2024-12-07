package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutDeviceRoute struct {
	deviceRepository            repositories.DeviceRepository
	bearerAuthRequestHeader     requests.BearerAuthRequestHeader
	logoutDeviceRequest         requests.LogoutDeviceRequest
	errorBadRequestResponse     responses.ErrorBadRequestResponse
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewLogoutDeviceRoute(
	deviceRepository repositories.DeviceRepository,
	bearerAuthRequestHeader requests.BearerAuthRequestHeader,
	logoutDeviceRequest requests.LogoutDeviceRequest,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *LogoutDeviceRoute {
	return &LogoutDeviceRoute{
		deviceRepository:            deviceRepository,
		bearerAuthRequestHeader:     bearerAuthRequestHeader,
		logoutDeviceRequest:         logoutDeviceRequest,
		errorBadRequestResponse:     errorBadRequestResponse,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver LogoutDeviceRoute) Method() string {
	return "POST"
}

func (receiver LogoutDeviceRoute) Pattern() string {
	return "/logout_device"
}

func (receiver LogoutDeviceRoute) Handle(context *gin.Context) {
	headers, err := receiver.bearerAuthRequestHeader.Make(context)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, err := receiver.logoutDeviceRequest.Make(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			receiver.errorBadRequestResponse.Make(err),
		)

		return
	}

	err = receiver.deviceRepository.DeleteDeviceByUUIDAndUserUUID(
		request.GetBody().GetDeviceUUID(),
		headers.GetAccessToken().GetSubject(),
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
