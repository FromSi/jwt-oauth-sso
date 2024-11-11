package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutRoute struct {
	config           configs.TokenConfig
	deviceRepository repositories.DeviceRepository
}

func NewLogoutRoute(
	config configs.TokenConfig,
	deviceRepository repositories.DeviceRepository,
) *LogoutRoute {
	return &LogoutRoute{
		config:           config,
		deviceRepository: deviceRepository,
	}
}

func (receiver LogoutRoute) Method() string {
	return "POST"
}

func (receiver LogoutRoute) Pattern() string {
	return "/logout"
}

func (receiver LogoutRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = requests.NewLogoutRequest(context)

	err = receiver.deviceRepository.DeleteDeviceByUUID(headers.AccessToken.DeviceUUID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
