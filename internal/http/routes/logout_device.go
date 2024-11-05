package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutDeviceRoute struct {
	config           configs.TokenConfig
	deviceRepository repositories.DeviceRepository
}

func NewLogoutDeviceRoute(
	config configs.TokenConfig,
	deviceRepository repositories.DeviceRepository,
) *LogoutDeviceRoute {
	return &LogoutDeviceRoute{
		config:           config,
		deviceRepository: deviceRepository,
	}
}

func (receiver LogoutDeviceRoute) Method() string {
	return "POST"
}

func (receiver LogoutDeviceRoute) Pattern() string {
	return "/logout_device"
}

func (receiver LogoutDeviceRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, errResponse := requests.NewLogoutDeviceRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	err = receiver.deviceRepository.DeleteDeviceByUUIDAndUserUUID(request.Body.DeviceUUID, headers.AccessToken.Subject)

	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorInternalServerResponse(err))

		return
	}

	context.Status(http.StatusAccepted)
}
