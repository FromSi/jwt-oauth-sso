package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_bearer_auth_request_header.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests BearerAuthRequestHeader
type BearerAuthRequestHeader interface {
	Make(*gin.Context) (BearerAuthRequestHeader, error)
	GetAccessToken() tokens.AccessToken
}
