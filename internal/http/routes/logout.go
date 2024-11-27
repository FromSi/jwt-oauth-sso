package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutRoute struct {
	deviceRepository   repositories.DeviceRepository
	accessTokenBuilder tokens.AccessTokenBuilder
}

func NewLogoutRoute(
	deviceRepository repositories.DeviceRepository,
	accessTokenBuilder tokens.AccessTokenBuilder,
) *LogoutRoute {
	return &LogoutRoute{
		deviceRepository:   deviceRepository,
		accessTokenBuilder: accessTokenBuilder,
	}
}

func (receiver LogoutRoute) Method() string {
	return "POST"
}

func (receiver LogoutRoute) Pattern() string {
	return "/logout"
}

func (receiver LogoutRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.accessTokenBuilder)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = requests.NewLogoutRequest(context)

	err = receiver.deviceRepository.DeleteDeviceByUUID(headers.AccessToken.GetDeviceUUID())

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
