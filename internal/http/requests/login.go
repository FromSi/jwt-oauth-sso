package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_login_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LoginRequest
type LoginRequest interface {
	Make(*gin.Context) (LoginRequest, error)
	GetBody() LoginRequestBody
	GetIP() string
	GetUserAgent() string
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_login_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LoginRequestBody
type LoginRequestBody interface {
	Make(*gin.Context) (LoginRequestBody, error)
	GetEmail() string
	GetPassword() string
}
