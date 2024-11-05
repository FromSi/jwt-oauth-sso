package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type PasswordResetWithOldRequest struct {
	Body PasswordResetWithOldRequestBody
}

func NewPasswordResetWithOldRequest(context *gin.Context) (*PasswordResetWithOldRequest, *responses.ErrorBadRequestResponse) {
	var request PasswordResetWithOldRequest

	requestBody, err := NewPasswordResetWithOldRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type PasswordResetWithOldRequestBody struct {
	OldPassword string `json:"oldPassword" binding:"required,password"`
	NewPassword string `json:"newPassword" binding:"required,password"`
}

func NewPasswordResetWithOldRequestBody(context *gin.Context) (*PasswordResetWithOldRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody PasswordResetWithOldRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
