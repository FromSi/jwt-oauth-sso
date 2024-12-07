package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_password_reset_with_old_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests PasswordResetWithOldRequest
type PasswordResetWithOldRequest interface {
	Make(*gin.Context) (PasswordResetWithOldRequest, error)
	GetBody() PasswordResetWithOldRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_password_reset_with_old_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests PasswordResetWithOldRequestBody
type PasswordResetWithOldRequestBody interface {
	Make(*gin.Context) (PasswordResetWithOldRequestBody, error)
	GetOldPassword() string
	GetNewPassword() string
}
