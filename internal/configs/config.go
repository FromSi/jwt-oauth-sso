package configs

type AppConfig interface {
	GetName() string
	GetHost() string
	GetPort() int
}

type TokenConfig interface {
	GetIssuerName() string
	GetAudienceName() string
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
