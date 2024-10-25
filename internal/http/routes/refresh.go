package routes

import (
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshRoute struct {
}

func NewRefreshRoute() *RefreshRoute {
	return &RefreshRoute{}
}

func (receiver RefreshRoute) Method() string {
	return "POST"
}

func (receiver RefreshRoute) Pattern() string {
	return "/refresh"
}

func (receiver RefreshRoute) Handle(context *gin.Context) {
	_, err := requests.NewRefreshRequest(context)

	if err != nil {
		context.Status(http.StatusBadRequest)

		return
	}

	fmt.Println(map[string]any{})

	context.Status(http.StatusContinue)
}
