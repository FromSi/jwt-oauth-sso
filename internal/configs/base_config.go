package configs

import (
	"flag"
	"github.com/spf13/viper"
)

const (
	BaseConfigDefaultAppName                        = "jwt-oauth-sso"
	BaseConfigDefaultAppHost                        = "localhost"
	BaseConfigDefaultAppPort                        = 8080
	BaseConfigDefaultAppDebug                       = true
	BaseConfigDefaultTokenIssuerName                = "jwt-oauth-sso"
	BaseConfigDefaultTokenAudienceName              = "user"
	BaseConfigDefaultTokenExpirationResetInDays     = 1
	BaseConfigDefaultTokenExpirationRefreshInDays   = 30
	BaseConfigDefaultTokenExpirationAccessInMinutes = 30
	BaseConfigDefaultTokenSecretKey                 = "secret"
	BaseConfigDefaultDatabaseDsn                    = "file::memory:"
)

type BaseConfig struct {
	appName                        string
	appHost                        string
	appPort                        int
	appDebug                       bool
	tokenIssuerName                string
	tokenAudienceName              string
	tokenExpirationResetInDays     int
	tokenExpirationRefreshInDays   int
	tokenExpirationAccessInMinutes int
	tokenSecretKey                 string
	databaseDsn                    string
}

func NewBaseConfig() *BaseConfig {
	var config BaseConfig

	viper.SetDefault(
		"app.name",
		BaseConfigDefaultAppName,
	)

	viper.SetDefault(
		"app.host",
		BaseConfigDefaultAppHost,
	)

	viper.SetDefault(
		"app.port",
		BaseConfigDefaultAppPort,
	)

	viper.SetDefault(
		"app.debug",
		BaseConfigDefaultAppDebug,
	)

	viper.SetDefault(
		"token.issuer_name",
		BaseConfigDefaultTokenIssuerName,
	)

	viper.SetDefault(
		"token.audience_name",
		BaseConfigDefaultTokenAudienceName,
	)

	viper.SetDefault(
		"token.expiration_reset_in_days",
		BaseConfigDefaultTokenExpirationResetInDays,
	)

	viper.SetDefault(
		"token.expiration_refresh_in_days",
		BaseConfigDefaultTokenExpirationRefreshInDays,
	)

	viper.SetDefault(
		"token.expiration_access_in_minutes",
		BaseConfigDefaultTokenExpirationAccessInMinutes,
	)

	viper.SetDefault(
		"token.secret_key",
		BaseConfigDefaultTokenSecretKey,
	)

	viper.SetDefault(
		"database.dsn", BaseConfigDefaultDatabaseDsn,
	)

	config.appName = viper.
		GetString("app.name")

	config.appHost = viper.
		GetString("app.host")

	config.appPort = viper.
		GetInt("app.port")

	config.appDebug = viper.
		GetBool("app.debug")

	config.tokenIssuerName = viper.
		GetString("token.issuer_name")

	config.tokenAudienceName = viper.
		GetString("token.audience_name")

	config.tokenExpirationResetInDays = viper.
		GetInt("token.expiration_reset_in_days")

	config.tokenExpirationRefreshInDays = viper.
		GetInt("token.expiration_refresh_in_days")

	config.tokenExpirationAccessInMinutes = viper.
		GetInt("token.expiration_access_in_minutes")

	config.tokenSecretKey = viper.
		GetString("token.secret_key")

	config.databaseDsn = viper.
		GetString("database.dsn")

	return &config
}

func NewBaseConfigWithYamlFile() (*BaseConfig, error) {
	var dirPath, filename string

	flag.StringVar(
		&filename,
		"config_filename",
		"config",
		"configuration filename. e.g: config",
	)

	flag.StringVar(
		&dirPath,
		"config_dir_path",
		".",
		"configuration file directory path",
	)

	flag.Parse()

	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dirPath)

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return NewBaseConfig(), nil
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

func (receiver BaseConfig) GetDebug() bool {
	return receiver.appDebug
}

func (receiver BaseConfig) GetIssuerName() string {
	return receiver.tokenIssuerName
}

func (receiver BaseConfig) GetAudienceName() string {
	return receiver.tokenAudienceName
}

func (receiver BaseConfig) GetExpirationResetInDays() int {
	return receiver.tokenExpirationResetInDays
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
