package requests

import "github.com/gin-gonic/gin"

type DevicesRequest struct {
	Body DevicesRequestBody
}

func NewDevicesRequest(context *gin.Context) (*DevicesRequest, error) {
	var request DevicesRequest

	requestBody, err := NewDevicesRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type DevicesRequestBody struct{}

func NewDevicesRequestBody(context *gin.Context) (*DevicesRequestBody, error) {
	var requestBody DevicesRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
