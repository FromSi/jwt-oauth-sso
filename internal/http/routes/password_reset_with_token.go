package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PasswordResetWithTokenRoute struct {
}

func NewPasswordResetWithTokenRoute() *PasswordResetWithTokenRoute {
	return &PasswordResetWithTokenRoute{}
}

func (receiver PasswordResetWithTokenRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithTokenRoute) Pattern() string {
	return "/password_reset_with_token"
}

func (receiver PasswordResetWithTokenRoute) Handle(context *gin.Context) {
	request, err := requests.NewPasswordResetWithTokenRequest(context)

	if err != nil {
		context.Status(http.StatusBadRequest)

		return
	}

	fmt.Println(map[string]any{
		"token":        request.Body.Token,
		"new_password": request.Body.NewPassword,
	})

	context.Status(http.StatusContinue)
}
