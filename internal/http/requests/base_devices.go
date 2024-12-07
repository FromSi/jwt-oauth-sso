package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseDevicesRequest struct {
	Body DevicesRequestBody
}

func NewBaseDevicesRequest(
	requestBody DevicesRequestBody,
) *BaseDevicesRequest {
	return &BaseDevicesRequest{
		Body: requestBody,
	}
}

func (receiver BaseDevicesRequest) Make(
	context *gin.Context,
) DevicesRequest {
	receiver.Body = receiver.Body.Make(context)

	return &receiver
}

func (receiver BaseDevicesRequest) GetBody() DevicesRequestBody {
	return receiver.Body
}

type BaseDevicesRequestBody struct{}

func NewBaseDevicesRequestBody() *BaseDevicesRequestBody {
	return &BaseDevicesRequestBody{}
}

func (receiver BaseDevicesRequestBody) Make(
	context *gin.Context,
) DevicesRequestBody {
	return &BaseDevicesRequestBody{}
}
