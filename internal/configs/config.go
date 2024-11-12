package configs

//go:generate mockgen -destination=../../mocks/configs/mock_config.go -package=configs_mocks github.com/fromsi/jwt-oauth-sso/internal/configs AppConfig
type AppConfig interface {
	GetName() string
	GetHost() string
	GetPort() int
	GetDebug() bool
}

//go:generate mockgen -destination=../../mocks/configs/mock_config.go -package=configs_mocks github.com/fromsi/jwt-oauth-sso/internal/configs TokenConfig
type TokenConfig interface {
	GetIssuerName() string
	GetAudienceName() string
	GetExpirationResetInDays() int
	GetExpirationRefreshInDays() int
	GetExpirationAccessInMinutes() int
	GetSecretKey() string
}

//go:generate mockgen -destination=../../mocks/configs/mock_config.go -package=configs_mocks github.com/fromsi/jwt-oauth-sso/internal/configs DatabaseConfig
type DatabaseConfig interface {
	GetDsn() string
}

//go:generate mockgen -destination=../../mocks/configs/mock_config.go -package=configs_mocks github.com/fromsi/jwt-oauth-sso/internal/configs Config
type Config interface {
	AppConfig
	TokenConfig
	DatabaseConfig
}
