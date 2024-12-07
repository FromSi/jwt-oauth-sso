package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseLogoutRequest struct {
	Body LogoutRequestBody
}

func NewBaseLogoutRequest(
	requestBody LogoutRequestBody,
) *BaseLogoutRequest {
	return &BaseLogoutRequest{
		Body: requestBody,
	}
}

func (receiver BaseLogoutRequest) Make(
	context *gin.Context,
) LogoutRequest {
	receiver.Body = receiver.Body.Make(context)

	return &receiver
}

func (receiver BaseLogoutRequest) GetBody() LogoutRequestBody {
	return receiver.Body
}

type BaseLogoutRequestBody struct{}

func NewBaseLogoutRequestBody() *BaseLogoutRequestBody {
	return &BaseLogoutRequestBody{}
}

func (receiver BaseLogoutRequestBody) Make(
	context *gin.Context,
) LogoutRequestBody {
	return &BaseLogoutRequestBody{}
}
