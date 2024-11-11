package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutAllRoute struct {
	config           configs.TokenConfig
	deviceRepository repositories.DeviceRepository
}

func NewLogoutAllRoute(
	config configs.TokenConfig,
	deviceRepository repositories.DeviceRepository,
) *LogoutAllRoute {
	return &LogoutAllRoute{
		config:           config,
		deviceRepository: deviceRepository,
	}
}

func (receiver LogoutAllRoute) Method() string {
	return "POST"
}

func (receiver LogoutAllRoute) Pattern() string {
	return "/logout_all"
}

func (receiver LogoutAllRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = requests.NewLogoutAllRequest(context)

	err = receiver.deviceRepository.DeleteAllDevicesByUserUUID(headers.AccessToken.Subject)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
