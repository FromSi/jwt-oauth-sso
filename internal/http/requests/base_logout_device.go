package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseLogoutDeviceRequest struct {
	Body LogoutDeviceRequestBody
}

func NewBaseLogoutDeviceRequest(
	requestBody LogoutDeviceRequestBody,
) *BaseLogoutDeviceRequest {
	return &BaseLogoutDeviceRequest{
		Body: requestBody,
	}
}

func (receiver BaseLogoutDeviceRequest) Make(
	context *gin.Context,
) (LogoutDeviceRequest, error) {
	requestBody, err := receiver.Body.Make(context)

	if err != nil {
		return nil, err
	}

	receiver.Body = requestBody

	return &receiver, nil
}

func (receiver BaseLogoutDeviceRequest) GetBody() LogoutDeviceRequestBody {
	return receiver.Body
}

type BaseLogoutDeviceRequestBody struct {
	DeviceUUID string `json:"deviceUuid" binding:"required,uuid4"`
}

func NewBaseLogoutDeviceRequestBody() *BaseLogoutDeviceRequestBody {
	return &BaseLogoutDeviceRequestBody{}
}

func (receiver BaseLogoutDeviceRequestBody) Make(
	context *gin.Context,
) (LogoutDeviceRequestBody, error) {
	var requestBody BaseLogoutDeviceRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BaseLogoutDeviceRequestBody) GetDeviceUUID() string {
	return receiver.DeviceUUID
}
