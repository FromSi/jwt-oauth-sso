package requests

import "github.com/gin-gonic/gin"

type PasswordResetWithOldRequest struct {
	Body PasswordResetWithOldRequestBody
}

func NewPasswordResetWithOldRequest(context *gin.Context) (*PasswordResetWithOldRequest, error) {
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

func NewPasswordResetWithOldRequestBody(context *gin.Context) (*PasswordResetWithOldRequestBody, error) {
	var requestBody PasswordResetWithOldRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
