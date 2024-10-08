package configs

import (
	"flag"
	"github.com/spf13/viper"
)

const (
	BaseConfigDefaultAppName                        = "jwt-oauth-sso"
	BaseConfigDefaultAppHost                        = "localhost"
	BaseConfigDefaultAppPort                        = 8080
	BaseConfigDefaultTokenIssuerName                = "jwt-oauth-sso"
	BaseConfigDefaultTokenAudienceName              = "user"
	BaseConfigDefaultTokenExpirationRefreshInDays   = 30
	BaseConfigDefaultTokenExpirationAccessInMinutes = 30
	BaseConfigDefaultTokenSecretKey                 = "secret"
	BaseConfigDefaultDatabaseDsn                    = "file::memory:"
)

type BaseConfig struct {
	appName                        string
	appHost                        string
	appPort                        int
	tokenIssuerName                string
	tokenAudienceName              string
	tokenExpirationRefreshInDays   int
	tokenExpirationAccessInMinutes int
	tokenSecretKey                 string
	databaseDsn                    string
}

func NewBaseConfig(onlyDefaultValues bool) (*BaseConfig, error) {
	var dirPath, filename string
	var err error

	if !onlyDefaultValues {
		flag.StringVar(&filename, "config_filename", "config", "configuration filename. e.g: config")
		flag.StringVar(&dirPath, "config_dir_path", ".", "configuration file directory path")
		flag.Parse()

		viper.SetConfigName(filename)
		viper.SetConfigType("yaml")
		viper.AddConfigPath(dirPath)

		err = viper.ReadInConfig()
	}

	var config BaseConfig

	viper.SetDefault("app.name", BaseConfigDefaultAppName)
	viper.SetDefault("app.host", BaseConfigDefaultAppHost)
	viper.SetDefault("app.port", BaseConfigDefaultAppPort)
	viper.SetDefault("token.issuer_name", BaseConfigDefaultTokenIssuerName)
	viper.SetDefault("token.audience_name", BaseConfigDefaultTokenAudienceName)
	viper.SetDefault("token.expiration_refresh_in_days", BaseConfigDefaultTokenExpirationRefreshInDays)
	viper.SetDefault("token.expiration_access_in_minutes", BaseConfigDefaultTokenExpirationAccessInMinutes)
	viper.SetDefault("token.secret_key", BaseConfigDefaultTokenSecretKey)
	viper.SetDefault("database.dsn", BaseConfigDefaultDatabaseDsn)

	config.appName = viper.GetString("app.name")
	config.appHost = viper.GetString("app.host")
	config.appPort = viper.GetInt("app.port")
	config.tokenIssuerName = viper.GetString("token.issuer_name")
	config.tokenAudienceName = viper.GetString("token.audience_name")
	config.tokenExpirationRefreshInDays = viper.GetInt("token.expiration_refresh_in_days")
	config.tokenExpirationAccessInMinutes = viper.GetInt("token.expiration_access_in_minutes")
	config.tokenSecretKey = viper.GetString("token.secret_key")
	config.databaseDsn = viper.GetString("database.dsn")

	return &config, err
}

func (receiver BaseConfig) GetName() string {
	return receiver.appName
}

func (receiver BaseConfig) GetHost() string {
	return receiver.appHost
}

func (receiver BaseConfig) GetPort() int {
	return receiver.appPort
}

func (receiver BaseConfig) GetIssuerName() string {
	return receiver.tokenIssuerName
}

func (receiver BaseConfig) GetAudienceName() string {
	return receiver.tokenAudienceName
}

func (receiver BaseConfig) GetExpirationRefreshInDays() int {
	return receiver.tokenExpirationRefreshInDays
}

func (receiver BaseConfig) GetExpirationAccessInMinutes() int {
	return receiver.tokenExpirationAccessInMinutes
}

func (receiver BaseConfig) GetSecretKey() string {
	return receiver.tokenSecretKey
}

func (receiver BaseConfig) GetDsn() string {
	return receiver.databaseDsn
}
