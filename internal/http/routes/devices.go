package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DevicesRoute struct {
	deviceRepository        repositories.DeviceRepository
	bearerAuthRequestHeader requests.BearerAuthRequestHeader
	devicesRequest          requests.DevicesRequest
	successDevicesResponse  responses.SuccessDevicesResponse
}

func NewDevicesRoute(
	deviceRepository repositories.DeviceRepository,
	bearerAuthRequestHeader requests.BearerAuthRequestHeader,
	devicesRequest requests.DevicesRequest,
	successDevicesResponse responses.SuccessDevicesResponse,
) *DevicesRoute {
	return &DevicesRoute{
		deviceRepository:        deviceRepository,
		bearerAuthRequestHeader: bearerAuthRequestHeader,
		devicesRequest:          devicesRequest,
		successDevicesResponse:  successDevicesResponse,
	}
}

func (receiver DevicesRoute) Method() string {
	return "GET"
}

func (receiver DevicesRoute) Pattern() string {
	return "/devices"
}

func (receiver DevicesRoute) Handle(context *gin.Context) {
	headers, err := receiver.bearerAuthRequestHeader.Make(context)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = receiver.devicesRequest.Make(context)

	devices := receiver.deviceRepository.GetDevicesByUserUUID(headers.GetAccessToken().GetSubject())

	response := receiver.successDevicesResponse.Make(devices)

	context.JSON(http.StatusOK, response)
}
