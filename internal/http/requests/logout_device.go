package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type LogoutDeviceRequest struct {
	Body LogoutDeviceRequestBody
}

func NewLogoutDeviceRequest(
	context *gin.Context,
) (*LogoutDeviceRequest, *responses.ErrorBadRequestResponse) {
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

func NewLogoutDeviceRequestBody(
	context *gin.Context,
) (*LogoutDeviceRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody LogoutDeviceRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
