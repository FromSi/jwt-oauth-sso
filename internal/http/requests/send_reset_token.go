package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_send_reset_token_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests SendResetTokenRequest
type SendResetTokenRequest interface {
	Make(*gin.Context) (SendResetTokenRequest, error)
	GetBody() SendResetTokenRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_send_reset_token_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests SendResetTokenRequestBody
type SendResetTokenRequestBody interface {
	Make(*gin.Context) (SendResetTokenRequestBody, error)
	GetEmail() string
}
