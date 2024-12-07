package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_refresh_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests RefreshRequest
type RefreshRequest interface {
	Make(*gin.Context) (RefreshRequest, error)
	GetBody() RefreshRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_refresh_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests RefreshRequestBody
type RefreshRequestBody interface {
	Make(*gin.Context) (RefreshRequestBody, error)
	GetRefreshToken() string
}
