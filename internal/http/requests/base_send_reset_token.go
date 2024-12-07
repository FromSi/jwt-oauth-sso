package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseSendResetTokenRequest struct {
	Body SendResetTokenRequestBody
}

func NewBaseSendResetTokenRequest(
	requestBody SendResetTokenRequestBody,
) *BaseSendResetTokenRequest {
	return &BaseSendResetTokenRequest{
		Body: requestBody,
	}
}

func (receiver BaseSendResetTokenRequest) Make(
	context *gin.Context,
) (SendResetTokenRequest, error) {
	requestBody, err := receiver.Body.Make(context)

	if err != nil {
		return nil, err
	}

	receiver.Body = requestBody

	return &receiver, nil
}

func (receiver BaseSendResetTokenRequest) GetBody() SendResetTokenRequestBody {
	return receiver.Body
}

type BaseSendResetTokenRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

func NewBaseSendResetTokenRequestBody() *BaseSendResetTokenRequestBody {
	return &BaseSendResetTokenRequestBody{}
}

func (receiver BaseSendResetTokenRequestBody) Make(
	context *gin.Context,
) (SendResetTokenRequestBody, error) {
	var requestBody BaseSendResetTokenRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BaseSendResetTokenRequestBody) GetEmail() string {
	return receiver.Email
}
