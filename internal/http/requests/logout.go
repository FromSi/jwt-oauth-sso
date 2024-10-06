package requests

import "github.com/gin-gonic/gin"

type LogoutRequest struct {
	Body LogoutRequestBody
}

func NewLogoutRequest(context *gin.Context) (*LogoutRequest, error) {
	var request LogoutRequest

	requestBody, err := NewLogoutRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type LogoutRequestBody struct{}

func NewLogoutRequestBody(context *gin.Context) (*LogoutRequestBody, error) {
	var requestBody LogoutRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
