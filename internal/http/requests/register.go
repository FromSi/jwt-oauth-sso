package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Body RegisterRequestBody
}

func NewRegisterRequest(context *gin.Context) (*RegisterRequest, *responses.ErrorBadRequestResponse) {
	var request RegisterRequest

	requestBody, err := NewRegisterRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type RegisterRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func NewRegisterRequestBody(context *gin.Context) (*RegisterRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody RegisterRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
