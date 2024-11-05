package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Body      RegisterRequestBody
	IP        string
	UserAgent string
}

func NewRegisterRequest(context *gin.Context) (*RegisterRequest, *responses.ErrorBadRequestResponse) {
	var request RegisterRequest

	requestBody, err := NewRegisterRequestBody(context)

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

type RegisterRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func NewRegisterRequestBody(context *gin.Context) (*RegisterRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody RegisterRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
