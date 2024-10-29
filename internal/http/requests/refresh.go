package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type RefreshRequest struct {
	Body RefreshRequestBody
}

func NewRefreshRequest(context *gin.Context) (*RefreshRequest, *responses.ErrorBadRequestResponse) {
	var request RefreshRequest

	requestBody, err := NewRefreshRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type RefreshRequestBody struct {
	RefreshToken string `json:"refreshToken" binding:"required,uuid4"`
}

func NewRefreshRequestBody(context *gin.Context) (*RefreshRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody RefreshRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
