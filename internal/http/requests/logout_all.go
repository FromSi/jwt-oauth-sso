package requests

import (
	"github.com/gin-gonic/gin"
)

type LogoutAllRequest struct {
	Body LogoutAllRequestBody
}

func NewLogoutAllRequest(context *gin.Context) *LogoutAllRequest {
	var request LogoutAllRequest

	requestBody := NewLogoutAllRequestBody(context)

	request.Body = *requestBody

	return &request
}

type LogoutAllRequestBody struct{}

func NewLogoutAllRequestBody(
	context *gin.Context,
) *LogoutAllRequestBody {
	var requestBody LogoutAllRequestBody

	return &requestBody
}
