package main

import (
	"context"
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/http/routes"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/fromsi/jwt-oauth-sso/internal/validator_rules"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type FxParams struct {
	fx.In

	Config configs.Config
	Routes []routes.Route `group:"routes"`
}

func main() {
	err := validator_rules.BindPassword()

	if err != nil {
		panic(err)
	}

	fx.New(CreateApp()).Run()
}

func CreateApp() fx.Option {
	routeAnnotationGroup := fx.ResultTags(`group:"routes"`)

	return fx.Options(
		fx.Provide(
			NewHTTPServer,
			NewDatabaseConnection,

			fx.Annotate(
				configs.NewBaseConfigWithYamlFile,
				fx.As(new(configs.Config)),
				fx.As(new(configs.AppConfig)),
				fx.As(new(configs.TokenConfig)),
				fx.As(new(configs.DatabaseConfig)),
			),

			fx.Annotate(
				tokens.NewJwtAccessTokenBuilder,
				fx.As(new(tokens.AccessTokenBuilder)),
			),

			fx.Annotate(
				repositories.NewBaseUserBuilder,
				fx.As(new(repositories.UserBuilder)),
			),
			fx.Annotate(
				repositories.NewBaseResetTokenBuilder,
				fx.As(new(repositories.ResetTokenBuilder)),
			),
			fx.Annotate(
				repositories.NewBaseDeviceBuilder,
				fx.As(new(repositories.DeviceBuilder)),
			),

			fx.Annotate(
				repositories.NewGormDeviceRepository,
				fx.As(new(repositories.DeviceRepository)),
				fx.As(new(repositories.QueryDeviceRepository)),
				fx.As(new(repositories.MutableDeviceRepository)),
			),
			fx.Annotate(
				repositories.NewGormResetTokenRepository,
				fx.As(new(repositories.ResetTokenRepository)),
				fx.As(new(repositories.QueryResetTokenRepository)),
				fx.As(new(repositories.MutableResetTokenRepository)),
			),
			fx.Annotate(
				repositories.NewGormUserRepository,
				fx.As(new(repositories.UserRepository)),
				fx.As(new(repositories.QueryUserRepository)),
				fx.As(new(repositories.MutableUserRepository)),
			),

			fx.Annotate(
				requests.NewBaseBearerAuthRequestHeader,
				fx.As(new(requests.BearerAuthRequestHeader)),
			),
			fx.Annotate(
				requests.NewBaseDevicesRequest,
				fx.As(new(requests.DevicesRequest)),
			),
			fx.Annotate(
				requests.NewBaseDevicesRequestBody,
				fx.As(new(requests.DevicesRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseLoginRequest,
				fx.As(new(requests.LoginRequest)),
			),
			fx.Annotate(
				requests.NewBaseLoginRequestBody,
				fx.As(new(requests.LoginRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseLogoutRequest,
				fx.As(new(requests.LogoutRequest)),
			),
			fx.Annotate(
				requests.NewBaseLogoutRequestBody,
				fx.As(new(requests.LogoutRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseLogoutAllRequest,
				fx.As(new(requests.LogoutAllRequest)),
			),
			fx.Annotate(
				requests.NewBaseLogoutAllRequestBody,
				fx.As(new(requests.LogoutAllRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseLogoutDeviceRequest,
				fx.As(new(requests.LogoutDeviceRequest)),
			),
			fx.Annotate(
				requests.NewBaseLogoutDeviceRequestBody,
				fx.As(new(requests.LogoutDeviceRequestBody)),
			),
			fx.Annotate(
				requests.NewBasePasswordResetWithOldRequest,
				fx.As(new(requests.PasswordResetWithOldRequest)),
			),
			fx.Annotate(
				requests.NewBasePasswordResetWithOldRequestBody,
				fx.As(new(requests.PasswordResetWithOldRequestBody)),
			),
			fx.Annotate(
				requests.NewBasePasswordResetWithTokenRequest,
				fx.As(new(requests.PasswordResetWithTokenRequest)),
			),
			fx.Annotate(
				requests.NewBasePasswordResetWithTokenRequestBody,
				fx.As(new(requests.PasswordResetWithTokenRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseRefreshRequest,
				fx.As(new(requests.RefreshRequest)),
			),
			fx.Annotate(
				requests.NewBaseRefreshRequestBody,
				fx.As(new(requests.RefreshRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseRegisterRequest,
				fx.As(new(requests.RegisterRequest)),
			),
			fx.Annotate(
				requests.NewBaseRegisterRequestBody,
				fx.As(new(requests.RegisterRequestBody)),
			),
			fx.Annotate(
				requests.NewBaseSendResetTokenRequest,
				fx.As(new(requests.SendResetTokenRequest)),
			),
			fx.Annotate(
				requests.NewBaseSendResetTokenRequestBody,
				fx.As(new(requests.SendResetTokenRequestBody)),
			),

			fx.Annotate(
				responses.NewBaseErrorBadRequestResponse,
				fx.As(new(responses.ErrorBadRequestResponse)),
			),
			fx.Annotate(
				responses.NewBaseErrorConflictResponse,
				fx.As(new(responses.ErrorConflictResponse)),
			),
			fx.Annotate(
				responses.NewBaseErrorInternalServerResponse,
				fx.As(new(responses.ErrorInternalServerResponse)),
			),
			fx.Annotate(
				responses.NewBaseSuccessDevicesResponse,
				fx.As(new(responses.SuccessDevicesResponse)),
			),
			fx.Annotate(
				responses.NewBaseSuccessLoginResponse,
				fx.As(new(responses.SuccessLoginResponse)),
			),
			fx.Annotate(
				responses.NewBaseSuccessRefreshResponse,
				fx.As(new(responses.SuccessRefreshResponse)),
			),
			fx.Annotate(
				responses.NewBaseSuccessRegisterResponse,
				fx.As(new(responses.SuccessRegisterResponse)),
			),

			fx.Annotate(
				services.NewBaseDeviceService,
				fx.As(new(services.DeviceService)),
				fx.As(new(services.QueryDeviceService)),
				fx.As(new(services.MutableDeviceService)),
			),
			fx.Annotate(
				services.NewBaseResetTokenService,
				fx.As(new(services.ResetTokenService)),
				fx.As(new(services.QueryResetTokenService)),
				fx.As(new(services.MutableResetTokenService)),
			),
			fx.Annotate(
				services.NewBaseUserService,
				fx.As(new(services.UserService)),
				fx.As(new(services.QueryUserService)),
				fx.As(new(services.MutableUserService)),
			),
			fx.Annotate(
				services.NewLogNotificationService,
				fx.As(new(services.NotificationService)),
				fx.As(new(services.QueryNotificationService)),
				fx.As(new(services.MutableNotificationService)),
			),

			fx.Annotate(
				routes.NewDevicesRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewLoginRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewLogoutRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewLogoutAllRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewLogoutDeviceRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewPasswordResetWithOldRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewPasswordResetWithTokenRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewRefreshRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewRegisterRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
			fx.Annotate(
				routes.NewSendResetTokenRoute,
				fx.As(new(routes.Route)),
				routeAnnotationGroup,
			),
		),
		fx.Invoke(func(*http.Server) {}),
	)
}

func NewHTTPServer(lifecycle fx.Lifecycle, fxParams FxParams) *http.Server {
	if fxParams.Config.GetDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.Default()

	for _, appRoute := range fxParams.Routes {
		route.Handle(appRoute.Method(), "/auth"+appRoute.Pattern(), appRoute.Handle)
	}

	server := &http.Server{
		Addr:    fxParams.Config.GetHost() + ":" + strconv.Itoa(fxParams.Config.GetPort()),
		Handler: route.Handler(),
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			fmt.Println("Starting HTTP server at", server.Addr)

			go func() {
				if err := server.ListenAndServe(); err != nil {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			return nil
		},
		OnStop: func(context context.Context) error {
			return server.Shutdown(context)
		},
	})

	return server
}

func NewDatabaseConnection(config configs.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.GetDsn()))

	if err != nil {
		panic(err)
	}

	return db
}
