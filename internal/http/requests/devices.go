package requests

import (
	"github.com/gin-gonic/gin"
)

type DevicesRequest struct {
	Body DevicesRequestBody
}

func NewDevicesRequest(context *gin.Context) *DevicesRequest {
	var request DevicesRequest

	requestBody := NewDevicesRequestBody(context)

	request.Body = *requestBody

	return &request
}

type DevicesRequestBody struct{}

func NewDevicesRequestBody(context *gin.Context) *DevicesRequestBody {
	var requestBody DevicesRequestBody

	return &requestBody
}
