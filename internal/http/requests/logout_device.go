package requests

import "github.com/gin-gonic/gin"

type LogoutDeviceRequest struct {
	Body LogoutDeviceRequestBody
}

func NewLogoutDeviceRequest(context *gin.Context) (*LogoutDeviceRequest, error) {
	var request LogoutDeviceRequest

	requestBody, err := NewLogoutDeviceRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type LogoutDeviceRequestBody struct {
	DeviceUUID string `json:"deviceUuid" binding:"required,uuid4"`
}

func NewLogoutDeviceRequestBody(context *gin.Context) (*LogoutDeviceRequestBody, error) {
	var requestBody LogoutDeviceRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
