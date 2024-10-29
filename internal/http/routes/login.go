package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRoute struct {
}

func NewLoginRoute() *LoginRoute {
	return &LoginRoute{}
}

func (receiver LoginRoute) Method() string {
	return "POST"
}

func (receiver LoginRoute) Pattern() string {
	return "/login"
}

func (receiver LoginRoute) Handle(context *gin.Context) {
	request, errResponse := requests.NewLoginRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	fmt.Println(map[string]any{
		"email":    request.Body.Email,
		"password": request.Body.Password,
	})

	context.Status(http.StatusContinue)
}