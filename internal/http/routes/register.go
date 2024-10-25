package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRoute struct {
}

func NewRegisterRoute() *RegisterRoute {
	return &RegisterRoute{}
}

func (receiver RegisterRoute) Method() string {
	return "POST"
}

func (receiver RegisterRoute) Pattern() string {
	return "/register"
}

func (receiver RegisterRoute) Handle(context *gin.Context) {
	request, err := requests.NewRegisterRequest(context)

	if err != nil {
		context.Status(http.StatusBadRequest)

		return
	}

	fmt.Println(map[string]any{
		"email":    request.Body.Email,
		"password": request.Body.Password,
	})

	context.Status(http.StatusContinue)
}
