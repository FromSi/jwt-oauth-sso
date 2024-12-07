package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_logout_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutRequest
type LogoutRequest interface {
	Make(*gin.Context) LogoutRequest
	GetBody() LogoutRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_logout_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutRequestBody
type LogoutRequestBody interface {
	Make(*gin.Context) LogoutRequestBody
}
