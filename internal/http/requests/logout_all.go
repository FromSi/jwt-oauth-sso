package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type LogoutAllRequest struct {
	Body LogoutAllRequestBody
}

func NewLogoutAllRequest(context *gin.Context) (*LogoutAllRequest, *responses.ErrorBadRequestResponse) {
	var request LogoutAllRequest

	requestBody, err := NewLogoutAllRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type LogoutAllRequestBody struct{}

func NewLogoutAllRequestBody(context *gin.Context) (*LogoutAllRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody LogoutAllRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
