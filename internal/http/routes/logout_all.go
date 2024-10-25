package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutAllRoute struct {
	config configs.TokenConfig
}

func NewLogoutAllRoute(config configs.TokenConfig) *LogoutAllRoute {
	return &LogoutAllRoute{
		config: config,
	}
}

func (receiver LogoutAllRoute) Method() string {
	return "POST"
}

func (receiver LogoutAllRoute) Pattern() string {
	return "/logout_all"
}

func (receiver LogoutAllRoute) Handle(context *gin.Context) {
	_, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_, err = requests.NewLogoutAllRequest(context)

	if err != nil {
		context.Status(http.StatusBadRequest)

		return
	}

	fmt.Println(map[string]any{})

	context.Status(http.StatusContinue)
}
