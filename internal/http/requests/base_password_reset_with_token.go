package requests

import (
	"github.com/gin-gonic/gin"
)

type BasePasswordResetWithTokenRequest struct {
	Body PasswordResetWithTokenRequestBody
}

func NewBasePasswordResetWithTokenRequest(
	requestBody PasswordResetWithTokenRequestBody,
) *BasePasswordResetWithTokenRequest {
	return &BasePasswordResetWithTokenRequest{
		Body: requestBody,
	}
}

func (receiver BasePasswordResetWithTokenRequest) Make(
	context *gin.Context,
) (PasswordResetWithTokenRequest, error) {
	requestBody, err := receiver.Body.Make(context)

	if err != nil {
		return nil, err
	}

	receiver.Body = requestBody

	return &receiver, nil
}

func (receiver BasePasswordResetWithTokenRequest) GetBody() PasswordResetWithTokenRequestBody {
	return receiver.Body
}

type BasePasswordResetWithTokenRequestBody struct {
	Token       string `json:"token" binding:"required,uuid4"`
	NewPassword string `json:"newPassword" binding:"required,password"`
}

func NewBasePasswordResetWithTokenRequestBody() *BasePasswordResetWithTokenRequestBody {
	return &BasePasswordResetWithTokenRequestBody{}
}

func (receiver BasePasswordResetWithTokenRequestBody) Make(
	context *gin.Context,
) (PasswordResetWithTokenRequestBody, error) {
	var requestBody BasePasswordResetWithTokenRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BasePasswordResetWithTokenRequestBody) GetToken() string {
	return receiver.Token
}

func (receiver BasePasswordResetWithTokenRequestBody) GetNewPassword() string {
	return receiver.NewPassword
}
