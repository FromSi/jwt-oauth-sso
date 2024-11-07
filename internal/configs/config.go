package configs

type AppConfig interface {
	GetName() string
	GetHost() string
	GetPort() int
	GetDebug() bool
}

type TokenConfig interface {
	GetIssuerName() string
	GetAudienceName() string
	GetExpirationResetInDays() int
	GetExpirationRefreshInDays() int
	GetExpirationAccessInMinutes() int
	GetSecretKey() string
}

type DatabaseConfig interface {
	GetDsn() string
}

type Config interface {
	AppConfig
	TokenConfig
	DatabaseConfig
}
