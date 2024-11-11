package requests

import (
	"github.com/gin-gonic/gin"
)

type LogoutRequest struct {
	Body LogoutRequestBody
}

func NewLogoutRequest(context *gin.Context) *LogoutRequest {
	var request LogoutRequest

	requestBody := NewLogoutRequestBody(context)

	request.Body = *requestBody

	return &request
}

type LogoutRequestBody struct{}

func NewLogoutRequestBody(context *gin.Context) *LogoutRequestBody {
	var requestBody LogoutRequestBody

	return &requestBody
}
