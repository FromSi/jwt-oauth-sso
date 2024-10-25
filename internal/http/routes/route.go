package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Method() string
	Pattern() string
	Handle(context *gin.Context)
}
