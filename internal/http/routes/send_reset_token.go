package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendResetTokenRoute struct {
}

func NewSendResetTokenRoute() *SendResetTokenRoute {
	return &SendResetTokenRoute{}
}

func (receiver SendResetTokenRoute) Method() string {
	return "POST"
}

func (receiver SendResetTokenRoute) Pattern() string {
	return "/send_reset_token"
}

func (receiver SendResetTokenRoute) Handle(context *gin.Context) {
	request, err := requests.NewSendResetTokenRequest(context)

	if err != nil {
		context.Status(http.StatusBadRequest)

		return
	}

	fmt.Println(map[string]any{
		"user_uuid": request.Body.UserUUID,
	})

	context.Status(http.StatusContinue)
}
