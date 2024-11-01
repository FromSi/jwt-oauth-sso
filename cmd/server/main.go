package main

import (
	"context"
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/routes"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
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

			fx.Annotate(routes.NewDevicesRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewLoginRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewLogoutRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewLogoutAllRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewLogoutDeviceRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewPasswordResetWithOldRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewPasswordResetWithTokenRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewRefreshRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewRegisterRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
			fx.Annotate(routes.NewSendResetTokenRoute, fx.As(new(routes.Route)), routeAnnotationGroup),
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
