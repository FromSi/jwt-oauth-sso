package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutDeviceRoute struct {
	config configs.TokenConfig
}

func NewLogoutDeviceRoute(config configs.TokenConfig) *LogoutDeviceRoute {
	return &LogoutDeviceRoute{
		config: config,
	}
}

func (receiver LogoutDeviceRoute) Method() string {
	return "POST"
}

func (receiver LogoutDeviceRoute) Pattern() string {
	return "/logout_device"
}

func (receiver LogoutDeviceRoute) Handle(context *gin.Context) {
	_, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, errResponse := requests.NewLogoutDeviceRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	fmt.Println(map[string]any{
		"device_uuid": request.Body.DeviceUUID,
	})

	context.Status(http.StatusContinue)
}
