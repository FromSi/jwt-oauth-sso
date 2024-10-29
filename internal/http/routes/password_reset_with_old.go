package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PasswordResetWithOldRoute struct {
	config configs.TokenConfig
}

func NewPasswordResetWithOldRoute(config configs.TokenConfig) *PasswordResetWithOldRoute {
	return &PasswordResetWithOldRoute{
		config: config,
	}
}

func (receiver PasswordResetWithOldRoute) Method() string {
	return "POST"
}

func (receiver PasswordResetWithOldRoute) Pattern() string {
	return "/password_reset_with_old"
}

func (receiver PasswordResetWithOldRoute) Handle(context *gin.Context) {
	_, err := requests.NewBearerAuthRequestHeader(context, receiver.config)

	if err != nil {
		context.Status(http.StatusUnauthorized)

		return
	}

	request, errResponse := requests.NewPasswordResetWithOldRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	fmt.Println(map[string]any{
		"old_password": request.Body.OldPassword,
		"new_password": request.Body.NewPassword,
	})

	context.Status(http.StatusContinue)
}
