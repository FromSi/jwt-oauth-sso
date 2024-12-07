package requests

import (
	"github.com/gin-gonic/gin"
)

type BaseLoginRequest struct {
	Body      LoginRequestBody
	IP        string
	UserAgent string
}

func NewBaseLoginRequest(
	requestBody LoginRequestBody,
) *BaseLoginRequest {
	return &BaseLoginRequest{
		Body: requestBody,
	}
}

func (receiver BaseLoginRequest) Make(
	context *gin.Context,
) (LoginRequest, error) {
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

func (receiver BaseLoginRequest) GetBody() LoginRequestBody {
	return receiver.Body
}

func (receiver BaseLoginRequest) GetIP() string {
	return receiver.IP
}

func (receiver BaseLoginRequest) GetUserAgent() string {
	return receiver.UserAgent
}

type BaseLoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func NewBaseLoginRequestBody() *BaseLoginRequestBody {
	return &BaseLoginRequestBody{}
}

func (receiver BaseLoginRequestBody) Make(
	context *gin.Context,
) (LoginRequestBody, error) {
	var requestBody BaseLoginRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return &requestBody, nil
}

func (receiver BaseLoginRequestBody) GetEmail() string {
	return receiver.Email
}

func (receiver BaseLoginRequestBody) GetPassword() string {
	return receiver.Password
}
