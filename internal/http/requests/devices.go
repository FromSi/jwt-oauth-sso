package requests

import (
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination=../../../mocks/http/requests/mock_devices_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests DevicesRequest
type DevicesRequest interface {
	Make(*gin.Context) DevicesRequest
	GetBody() DevicesRequestBody
}

//go:generate mockgen -destination=../../../mocks/http/requests/mock_devices_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests DevicesRequestBody
type DevicesRequestBody interface {
	Make(*gin.Context) DevicesRequestBody
}
