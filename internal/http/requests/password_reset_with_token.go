package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type PasswordResetWithTokenRequest struct {
	Body PasswordResetWithTokenRequestBody
}

func NewPasswordResetWithTokenRequest(context *gin.Context) (*PasswordResetWithTokenRequest, *responses.ErrorBadRequestResponse) {
	var request PasswordResetWithTokenRequest

	requestBody, err := NewPasswordResetWithTokenRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type PasswordResetWithTokenRequestBody struct {
	Token       string `json:"token" binding:"required,uuid4"`
	NewPassword string `json:"newPassword" binding:"required,password"`
}

func NewPasswordResetWithTokenRequestBody(context *gin.Context) (*PasswordResetWithTokenRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody PasswordResetWithTokenRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
