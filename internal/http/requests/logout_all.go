package requests

import "github.com/gin-gonic/gin"

type LogoutAllRequest struct {
	Body LogoutAllRequestBody
}

func NewLogoutAllRequest(context *gin.Context) (*LogoutAllRequest, error) {
	var request LogoutAllRequest

	requestBody, err := NewLogoutAllRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type LogoutAllRequestBody struct{}

func NewLogoutAllRequestBody(context *gin.Context) (*LogoutAllRequestBody, error) {
	var requestBody LogoutAllRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
