package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_logout_device_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutDeviceRequest
type LogoutDeviceRequest interface {
	Make(*gin.Context) (LogoutDeviceRequest, error)
	GetBody() LogoutDeviceRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_logout_device_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutDeviceRequestBody
type LogoutDeviceRequestBody interface {
	Make(*gin.Context) (LogoutDeviceRequestBody, error)
	GetDeviceUUID() string
}
