package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigSystem := make(chan os.Signal, 1)

	signal.Notify(sigSystem, syscall.SIGINT, syscall.SIGTERM)

	route := gin.Default()

	route.POST("/auth/login", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.POST("/auth/register", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.POST("/auth/refresh", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.POST("/auth/logout", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.POST("/auth/logout_all", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.POST("/auth/logout_device", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.GET("/auth/devices", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]string{"message": "Hello"})
	})

	route.POST("/auth/send_reset_token", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	route.POST("/auth/reset_password", func(context *gin.Context) {
		context.Status(http.StatusContinue)
	})

	go func() {
		route.Run(":8080")
	}()

	<-sigSystem
}
