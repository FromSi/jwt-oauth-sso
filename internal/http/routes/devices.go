package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DevicesRoute struct {
	config configs.TokenConfig
}

func NewDevicesRoute(config configs.TokenConfig) *DevicesRoute {
	return &DevicesRoute{
		config: config,
	}
}

func (receiver DevicesRoute) Method() string {
	return "GET"
}

func (receiver DevicesRoute) Pattern() string {
	return "/devices"
}

func (receiver DevicesRoute) Handle(context *gin.Context) {
	_, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_, errResponse := requests.NewDevicesRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	fmt.Println(map[string]any{})

	context.JSON(http.StatusOK, map[string]string{"message": "Hello"})
}
