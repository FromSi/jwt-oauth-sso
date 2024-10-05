package configs

import (
	"flag"
	"github.com/spf13/viper"
)

const (
	BaseConfigDefaultAppName                        = "jwt-oauth-sso"
	BaseConfigDefaultAppHost                        = "localhost"
	BaseConfigDefaultAppPort                        = 8080
	BaseConfigDefaultTokenExpirationRefreshInDays   = 30
	BaseConfigDefaultTokenExpirationAccessInMinutes = 30
	BaseConfigDefaultTokenSecretKey                 = "secret"
	BaseConfigDefaultDatabaseDsn                    = "file::memory:"
)

type BaseConfig struct {
	AppName                        string
	AppHost                        string
	AppPort                        int
	TokenExpirationRefreshInDays   int
	TokenExpirationAccessInMinutes int
	TokenSecretKey                 string
	DatabaseDsn                    string
}

func NewBaseConfig(onlyDefaultValues bool) (*BaseConfig, error) {
	var dirPath, filename string

	if !onlyDefaultValues {
		flag.StringVar(&filename, "config_filename", "config", "configuration filename. e.g: config")
		flag.StringVar(&dirPath, "config_dir_path", ".", "configuration file directory path")
		flag.Parse()

		viper.SetConfigName(filename)
		viper.SetConfigType("yaml")
		viper.AddConfigPath(dirPath)
		viper.ReadInConfig()
	}

	var config BaseConfig

	viper.SetDefault("app.name", BaseConfigDefaultAppName)
	viper.SetDefault("app.host", BaseConfigDefaultAppHost)
	viper.SetDefault("app.port", BaseConfigDefaultAppPort)
	viper.SetDefault("token.expiration_refresh_in_days", BaseConfigDefaultTokenExpirationRefreshInDays)
	viper.SetDefault("token.expiration_access_in_minutes", BaseConfigDefaultTokenExpirationAccessInMinutes)
	viper.SetDefault("token.secret_key", BaseConfigDefaultTokenSecretKey)
	viper.SetDefault("database.dsn", BaseConfigDefaultDatabaseDsn)

	config.AppName = viper.GetString("app.name")
	config.AppHost = viper.GetString("app.host")
	config.AppPort = viper.GetInt("app.port")
	config.TokenExpirationRefreshInDays = viper.GetInt("token.expiration_refresh_in_days")
	config.TokenExpirationAccessInMinutes = viper.GetInt("token.expiration_access_in_minutes")
	config.TokenSecretKey = viper.GetString("token.secret_key")
	config.DatabaseDsn = viper.GetString("database.dsn")

	return &config, nil
}

func (receiver BaseConfig) GetName() string {
	return receiver.AppName
}

func (receiver BaseConfig) GetHost() string {
	return receiver.AppHost
}

func (receiver BaseConfig) GetPort() int {
	return receiver.AppPort
}

func (receiver BaseConfig) GetExpirationRefreshInDays() int {
	return receiver.TokenExpirationRefreshInDays
}

func (receiver BaseConfig) GetExpirationAccessInMinutes() int {
	return receiver.TokenExpirationAccessInMinutes
}

func (receiver BaseConfig) GetSecretKey() string {
	return receiver.TokenSecretKey
}

func (receiver BaseConfig) GetDsn() string {
	return receiver.DatabaseDsn
}
