package main

import (
	"context"
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/validator_rules"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	config, err := configs.NewBaseConfig(false)

	if err != nil {
		panic(err)
	}

	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("password", validator_rules.Password)

		if err != nil {
			panic(err)
		}
	}

	route.POST("/auth/login", func(ctx *gin.Context) {
		request, err := requests.NewLoginRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{
			"email":    request.Body.Email,
			"password": request.Body.Password,
		})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/register", func(ctx *gin.Context) {
		request, err := requests.NewRegisterRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{
			"email":    request.Body.Email,
			"password": request.Body.Password,
		})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/refresh", func(ctx *gin.Context) {
		_, err = requests.NewRefreshRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/logout", func(ctx *gin.Context) {
		_, err := requests.NewBearerAuthRequestHeader(ctx, config)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)

			return
		}

		_, err = requests.NewLogoutRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/logout_all", func(ctx *gin.Context) {
		_, err := requests.NewBearerAuthRequestHeader(ctx, config)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)

			return
		}

		_, err = requests.NewLogoutAllRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/logout_device", func(ctx *gin.Context) {
		_, err := requests.NewBearerAuthRequestHeader(ctx, config)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)

			return
		}

		request, err := requests.NewLogoutDeviceRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{
			"device_uuid": request.Body.DeviceUUID,
		})

		ctx.Status(http.StatusContinue)
	})

	route.GET("/auth/devices", func(ctx *gin.Context) {
		_, err := requests.NewBearerAuthRequestHeader(ctx, config)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)

			return
		}

		_, err = requests.NewDevicesRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{})

		ctx.JSON(http.StatusOK, map[string]string{"message": "Hello"})
	})

	route.POST("/auth/send_reset_token", func(ctx *gin.Context) {
		request, err := requests.NewSendResetTokenRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{
			"user_uuid": request.Body.UserUUID,
		})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/password_reset_with_token", func(ctx *gin.Context) {
		request, err := requests.NewPasswordResetWithTokenRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{
			"token":        request.Body.Token,
			"new_password": request.Body.NewPassword,
		})

		ctx.Status(http.StatusContinue)
	})

	route.POST("/auth/password_reset_with_old", func(ctx *gin.Context) {
		_, err := requests.NewBearerAuthRequestHeader(ctx, config)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)

			return
		}

		request, err := requests.NewPasswordResetWithOldRequest(ctx)

		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		fmt.Println(map[string]any{
			"old_password": request.Body.OldPassword,
			"new_password": request.Body.NewPassword,
		})

		ctx.Status(http.StatusContinue)
	})

	server := &http.Server{
		Addr:    config.GetHost() + ":" + strconv.Itoa(config.GetPort()),
		Handler: route.Handler(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
