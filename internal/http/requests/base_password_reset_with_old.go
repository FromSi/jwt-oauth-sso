package requests

import (
	"github.com/gin-gonic/gin"
)

type BasePasswordResetWithOldRequest struct {
	Body PasswordResetWithOldRequestBody
}

func NewBasePasswordResetWithOldRequest(
	requestBody PasswordResetWithOldRequestBody,
) *BasePasswordResetWithOldRequest {
	return &BasePasswordResetWithOldRequest{
		Body: requestBody,
	}
}

func (receiver BasePasswordResetWithOldRequest) Make(
	context *gin.Context,
) (PasswordResetWithOldRequest, error) {
	requestBody, err := receiver.Body.Make(context)

	if err != nil {
		return nil, err
	}

	receiver.Body = requestBody

	return &receiver, nil
}

func (receiver BasePasswordResetWithOldRequest) GetBody() PasswordResetWithOldRequestBody {
	return receiver.Body
}

type BasePasswordResetWithOldRequestBody struct {
	OldPassword string `json:"oldPassword" binding:"required,password"`
	NewPassword string `json:"newPassword" binding:"required,password"`
}

func NewBasePasswordResetWithOldRequestBody() *BasePasswordResetWithOldRequestBody {
	return &BasePasswordResetWithOldRequestBody{}
}

func (receiver BasePasswordResetWithOldRequestBody) Make(
	context *gin.Context,
) (PasswordResetWithOldRequestBody, error) {
	var requestBody BasePasswordResetWithOldRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BasePasswordResetWithOldRequestBody) GetOldPassword() string {
	return receiver.OldPassword
}

func (receiver BasePasswordResetWithOldRequestBody) GetNewPassword() string {
	return receiver.NewPassword
}
