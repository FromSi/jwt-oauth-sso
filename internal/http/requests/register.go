package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_register_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests RegisterRequest
type RegisterRequest interface {
	Make(*gin.Context) (RegisterRequest, error)
	GetBody() RegisterRequestBody
	GetIP() string
	GetUserAgent() string
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_register_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests RegisterRequestBody
type RegisterRequestBody interface {
	Make(*gin.Context) (RegisterRequestBody, error)
	GetEmail() string
	GetPassword() string
}
