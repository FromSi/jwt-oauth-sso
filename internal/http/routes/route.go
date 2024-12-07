package routes

import "github.com/gin-gonic/gin"

//go:generate mockgen -destination=../../../mocks/http/routes/mock_route.go -package=routes_mocks github.com/fromsi/jwt-oauth-sso/internal/http/routes Route
type Route interface {
	Method() string
	Pattern() string
	Handle(*gin.Context)
}
