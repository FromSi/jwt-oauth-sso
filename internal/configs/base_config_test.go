package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseConfig(t *testing.T) {
	_, err := NewBaseConfig(true)

	assert.NoError(t, err)
}

func TestBaseConfig_GetName(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetName(), BaseConfigDefaultAppName)
}

func TestBaseConfig_GetHost(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetHost(), BaseConfigDefaultAppHost)
}

func TestBaseConfig_GetPort(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetHost(), BaseConfigDefaultAppHost)
}

func TestBaseConfig_GetIssuerName(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetIssuerName(), BaseConfigDefaultTokenIssuerName)
}

func TestBaseConfig_GetAudienceName(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetAudienceName(), BaseConfigDefaultTokenAudienceName)
}

func TestBaseConfig_GetExpirationRefreshInDays(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetExpirationRefreshInDays(), BaseConfigDefaultTokenExpirationRefreshInDays)
}

func TestBaseConfig_GetExpirationAccessInMinutes(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetExpirationAccessInMinutes(), BaseConfigDefaultTokenExpirationAccessInMinutes)
}

func TestBaseConfig_GetSecretKey(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetSecretKey(), BaseConfigDefaultTokenSecretKey)
}

func TestBaseConfig_GetDsn(t *testing.T) {
	config, _ := NewBaseConfig(true)

	assert.Equal(t, config.GetDsn(), BaseConfigDefaultDatabaseDsn)
}
