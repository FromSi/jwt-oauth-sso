package requests

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Body LoginRequestBody
}

func NewLoginRequest(context *gin.Context) (*LoginRequest, error) {
	var request LoginRequest

	requestBody, err := NewLoginRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func NewLoginRequestBody(context *gin.Context) (*LoginRequestBody, error) {
	var requestBody LoginRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
