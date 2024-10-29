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
	request, errResponse := requests.NewSendResetTokenRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	fmt.Println(map[string]any{
		"user_uuid": request.Body.UserUUID,
	})

	context.Status(http.StatusContinue)
}
