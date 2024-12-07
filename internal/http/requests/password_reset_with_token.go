package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_password_reset_with_token_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests PasswordResetWithTokenRequest
type PasswordResetWithTokenRequest interface {
	Make(*gin.Context) (PasswordResetWithTokenRequest, error)
	GetBody() PasswordResetWithTokenRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_password_reset_with_token_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests PasswordResetWithTokenRequestBody
type PasswordResetWithTokenRequestBody interface {
	Make(*gin.Context) (PasswordResetWithTokenRequestBody, error)
	GetToken() string
	GetNewPassword() string
}
