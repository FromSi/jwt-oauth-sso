package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutDeviceRoute struct {
	deviceRepository   repositories.DeviceRepository
	accessTokenBuilder tokens.AccessTokenBuilder
}

func NewLogoutDeviceRoute(
	deviceRepository repositories.DeviceRepository,
	accessTokenBuilder tokens.AccessTokenBuilder,
) *LogoutDeviceRoute {
	return &LogoutDeviceRoute{
		deviceRepository:   deviceRepository,
		accessTokenBuilder: accessTokenBuilder,
	}
}

func (receiver LogoutDeviceRoute) Method() string {
	return "POST"
}

func (receiver LogoutDeviceRoute) Pattern() string {
	return "/logout_device"
}

func (receiver LogoutDeviceRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.accessTokenBuilder)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, errResponse := requests.NewLogoutDeviceRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	err = receiver.deviceRepository.DeleteDeviceByUUIDAndUserUUID(
		request.Body.DeviceUUID,
		headers.AccessToken.GetSubject(),
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
