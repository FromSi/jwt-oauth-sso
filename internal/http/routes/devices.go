package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DevicesRoute struct {
	deviceRepository   repositories.DeviceRepository
	accessTokenBuilder tokens.AccessTokenBuilder
}

func NewDevicesRoute(
	deviceRepository repositories.DeviceRepository,
	accessTokenBuilder tokens.AccessTokenBuilder,
) *DevicesRoute {
	return &DevicesRoute{
		deviceRepository:   deviceRepository,
		accessTokenBuilder: accessTokenBuilder,
	}
}

func (receiver DevicesRoute) Method() string {
	return "GET"
}

func (receiver DevicesRoute) Pattern() string {
	return "/devices"
}

func (receiver DevicesRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.accessTokenBuilder)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = requests.NewDevicesRequest(context)

	devices := receiver.deviceRepository.GetDevicesByUserUUID(headers.AccessToken.GetSubject())

	response := responses.NewSuccessDevicesResponse(devices)

	context.JSON(http.StatusOK, response)
}
