package requests

import "github.com/gin-gonic/gin"

type RegisterRequest struct {
	Body RegisterRequestBody
}

func NewRegisterRequest(context *gin.Context) (*RegisterRequest, error) {
	var request RegisterRequest

	requestBody, err := NewRegisterRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type RegisterRequestBody struct {
	Email    string `from:"email" binding:"required,email"`
	Password string `from:"password" binding:"required,password"`
}

func NewRegisterRequestBody(context *gin.Context) (*RegisterRequestBody, error) {
	var requestBody RegisterRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
