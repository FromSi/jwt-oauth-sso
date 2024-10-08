package requests

import (
	"github.com/gin-gonic/gin"
)

type RefreshRequest struct {
	Body RefreshRequestBody
}

func NewRefreshRequest(context *gin.Context) (*RefreshRequest, error) {
	var request RefreshRequest

	requestBody, err := NewRefreshRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type RefreshRequestBody struct{}

func NewRefreshRequestBody(context *gin.Context) (*RefreshRequestBody, error) {
	var requestBody RefreshRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
