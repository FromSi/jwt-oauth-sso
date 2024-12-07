package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseLogoutAllRequest struct {
	Body LogoutAllRequestBody
}

func NewBaseLogoutAllRequest(
	requestBody LogoutAllRequestBody,
) *BaseLogoutAllRequest {
	return &BaseLogoutAllRequest{
		Body: requestBody,
	}
}

func (receiver BaseLogoutAllRequest) Make(
	context *gin.Context,
) LogoutAllRequest {
	receiver.Body = receiver.Body.Make(context)

	return &receiver
}

func (receiver BaseLogoutAllRequest) GetBody() LogoutAllRequestBody {
	return receiver.Body
}

type BaseLogoutAllRequestBody struct{}

func NewBaseLogoutAllRequestBody() *BaseLogoutAllRequestBody {
	return &BaseLogoutAllRequestBody{}
}

func (receiver BaseLogoutAllRequestBody) Make(
	context *gin.Context,
) LogoutAllRequestBody {
	return &BaseLogoutAllRequestBody{}
}
