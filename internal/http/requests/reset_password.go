package requests

import "github.com/gin-gonic/gin"

type ResetPasswordRequest struct {
	Body ResetPasswordRequestBody
}

func NewResetPasswordRequest(context *gin.Context) (*ResetPasswordRequest, error) {
	var request ResetPasswordRequest

	requestBody, err := NewResetPasswordRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type ResetPasswordRequestBody struct {
	Token string `from:"token" binding:"required,uuid4"`
}

func NewResetPasswordRequestBody(context *gin.Context) (*ResetPasswordRequestBody, error) {
	var requestBody ResetPasswordRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}
