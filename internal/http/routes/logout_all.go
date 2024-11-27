package routes

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutAllRoute struct {
	deviceRepository   repositories.DeviceRepository
	accessTokenBuilder tokens.AccessTokenBuilder
}

func NewLogoutAllRoute(
	deviceRepository repositories.DeviceRepository,
	accessTokenBuilder tokens.AccessTokenBuilder,
) *LogoutAllRoute {
	return &LogoutAllRoute{
		deviceRepository:   deviceRepository,
		accessTokenBuilder: accessTokenBuilder,
	}
}

func (receiver LogoutAllRoute) Method() string {
	return "POST"
}

func (receiver LogoutAllRoute) Pattern() string {
	return "/logout_all"
}

func (receiver LogoutAllRoute) Handle(context *gin.Context) {
	headers, err := requests.NewBearerAuthRequestHeader(context, receiver.accessTokenBuilder)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_ = requests.NewLogoutAllRequest(context)

	err = receiver.deviceRepository.DeleteAllDevicesByUserUUID(headers.AccessToken.GetSubject())

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.Status(http.StatusAccepted)
}
