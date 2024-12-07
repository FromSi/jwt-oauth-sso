package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseRegisterRequest struct {
	Body      RegisterRequestBody
	IP        string
	UserAgent string
}

func NewBaseRegisterRequest(
	requestBody RegisterRequestBody,
) *BaseRegisterRequest {
	return &BaseRegisterRequest{
		Body: requestBody,
	}
}

func (receiver BaseRegisterRequest) Make(
	context *gin.Context,
) (RegisterRequest, error) {
	requestBody, err := receiver.Body.Make(context)

	if err != nil {
		return nil, err
	}

	receiver.Body = requestBody

	receiver.IP = context.ClientIP()

	if receiver.IP == "" {
		receiver.IP = context.GetHeader("X-Real-Ip")
	}

	receiver.UserAgent = context.Request.UserAgent()

	return &receiver, nil
}

func (receiver BaseRegisterRequest) GetBody() RegisterRequestBody {
	return receiver.Body
}

func (receiver BaseRegisterRequest) GetIP() string {
	return receiver.IP
}

func (receiver BaseRegisterRequest) GetUserAgent() string {
	return receiver.UserAgent
}

type BaseRegisterRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func NewBaseRegisterRequestBody() *BaseRegisterRequestBody {
	return &BaseRegisterRequestBody{}
}

func (receiver BaseRegisterRequestBody) Make(
	context *gin.Context,
) (RegisterRequestBody, error) {
	var requestBody BaseRegisterRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BaseRegisterRequestBody) GetEmail() string {
	return receiver.Email
}

func (receiver BaseRegisterRequestBody) GetPassword() string {
	return receiver.Password
}
