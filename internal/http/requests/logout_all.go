package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_logout_all_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutAllRequest
type LogoutAllRequest interface {
	Make(*gin.Context) LogoutAllRequest
	GetBody() LogoutAllRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_logout_all_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutAllRequestBody
type LogoutAllRequestBody interface {
	Make(*gin.Context) LogoutAllRequestBody
}
