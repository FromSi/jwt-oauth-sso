package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DevicesRoute struct {
	config           configs.TokenConfig
	deviceRepository repositories.DeviceRepository
}

func NewDevicesRoute(
	config configs.TokenConfig,
	deviceRepository repositories.DeviceRepository,
) *DevicesRoute {
	return &DevicesRoute{
		config:           config,
		deviceRepository: deviceRepository,
	}
}

func (receiver DevicesRoute) Method() string {
	return "GET"
}

func (receiver DevicesRoute) Pattern() string {
	return "/devices"
}

func (receiver DevicesRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_, errResponse := requests.NewDevicesRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	devices := receiver.deviceRepository.GetDevicesByUserUUID(headers.AccessToken.Subject)

	response := responses.NewSuccessDevicesResponse(devices)

	context.JSON(http.StatusOK, response)
}
