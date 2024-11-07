package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Body      LoginRequestBody
	IP        string
	UserAgent string
}

func NewLoginRequest(
	context *gin.Context,
) (*LoginRequest, *responses.ErrorBadRequestResponse) {
	var request LoginRequest

	requestBody, err := NewLoginRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	request.IP = context.ClientIP()

	if request.IP == "" {
		request.IP = context.GetHeader("X-Real-Ip")
	}

	request.UserAgent = context.Request.UserAgent()

	return &request, nil
}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func NewLoginRequestBody(
	context *gin.Context,
) (*LoginRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody LoginRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
