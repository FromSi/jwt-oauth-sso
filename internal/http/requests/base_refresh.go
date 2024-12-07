package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseRefreshRequest struct {
	Body RefreshRequestBody
}

func NewBaseRefreshRequest(
	requestBody RefreshRequestBody,
) *BaseRefreshRequest {
	return &BaseRefreshRequest{
		Body: requestBody,
	}
}

func (receiver BaseRefreshRequest) Make(
	context *gin.Context,
) (RefreshRequest, error) {
	requestBody, err := receiver.Body.Make(context)

	if err != nil {
		return nil, err
	}

	receiver.Body = requestBody

	return &receiver, nil
}

func (receiver BaseRefreshRequest) GetBody() RefreshRequestBody {
	return receiver.Body
}

type BaseRefreshRequestBody struct {
	RefreshToken string `json:"refreshToken" binding:"required,uuid4"`
}

func NewBaseRefreshRequestBody() *BaseRefreshRequestBody {
	return &BaseRefreshRequestBody{}
}

func (receiver BaseRefreshRequestBody) Make(
	context *gin.Context,
) (RefreshRequestBody, error) {
	var requestBody BaseRefreshRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BaseRefreshRequestBody) GetRefreshToken() string {
	return receiver.RefreshToken
}
