package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutRoute struct {
	config configs.TokenConfig
}

func NewLogoutRoute(config configs.TokenConfig) *LogoutRoute {
	return &LogoutRoute{
		config: config,
	}
}

func (receiver LogoutRoute) Method() string {
	return "POST"
}

func (receiver LogoutRoute) Pattern() string {
	return "/logout"
}

func (receiver LogoutRoute) Handle(context *gin.Context) {
	_, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	_, err = requests.NewLogoutRequest(context)

	if err != nil {
		context.Status(http.StatusBadRequest)

		return
	}

	fmt.Println(map[string]any{})

	context.Status(http.StatusContinue)
}
