package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type LogoutRequest struct {
	Body LogoutRequestBody
}

func NewLogoutRequest(context *gin.Context) (*LogoutRequest, *responses.ErrorBadRequestResponse) {
	var request LogoutRequest

	requestBody, err := NewLogoutRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type LogoutRequestBody struct{}

func NewLogoutRequestBody(context *gin.Context) (*LogoutRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody LogoutRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		println(err.Error())
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
