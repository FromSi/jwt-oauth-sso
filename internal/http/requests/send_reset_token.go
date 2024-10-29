package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type SendResetTokenRequest struct {
	Body SendResetTokenRequestBody
}

func NewSendResetTokenRequest(context *gin.Context) (*SendResetTokenRequest, *responses.ErrorBadRequestResponse) {
	var request SendResetTokenRequest

	requestBody, err := NewSendResetTokenRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type SendResetTokenRequestBody struct {
	UserUUID string `json:"userUuid" binding:"required,uuid4"`
}

func NewSendResetTokenRequestBody(context *gin.Context) (*SendResetTokenRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody SendResetTokenRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
