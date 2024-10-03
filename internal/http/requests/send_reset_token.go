package requests

import "github.com/gin-gonic/gin"

type SendResetTokenRequest struct {
	Body SendResetTokenRequestBody
}

func NewSendResetTokenRequest(context *gin.Context) (*SendResetTokenRequest, error) {
	var request SendResetTokenRequest

	requestBody, err := NewSendResetTokenRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type SendResetTokenRequestBody struct {
	UserUUID string `from:"user_uuid" binding:"required,uuid4"`
}

func NewSendResetTokenRequestBody(context *gin.Context) (*SendResetTokenRequestBody, error) {
	var requestBody SendResetTokenRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
