package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/gin-gonic/gin"
)

type DevicesRequest struct {
	Body DevicesRequestBody
}

func NewDevicesRequest(
	context *gin.Context,
) (*DevicesRequest, *responses.ErrorBadRequestResponse) {
	var request DevicesRequest

	requestBody, err := NewDevicesRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type DevicesRequestBody struct{}

func NewDevicesRequestBody(
	context *gin.Context,
) (*DevicesRequestBody, *responses.ErrorBadRequestResponse) {
	var requestBody DevicesRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, responses.NewErrorBadRequestResponseByError(err)
	}

	return &requestBody, nil
}
