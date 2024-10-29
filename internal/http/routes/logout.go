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

	_, errResponse := requests.NewLogoutRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	fmt.Println(map[string]any{})

	context.Status(http.StatusContinue)
}
