package tokens

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewAccessToken(t *testing.T) {
	config := configs.NewBaseConfig()

	currentTime := time.Now().Truncate(time.Second)

	expirationTime := currentTime.
		Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes()))

	token, err := NewAccessToken(
		config,
		"1",
		"1",
		"1",
		currentTime,
	)

	assert.NoError(t, err)

	assert.Equal(t, token.Issuer, config.GetIssuerName())
	assert.Equal(t, token.Audience, config.GetAudienceName())
	assert.Equal(t, token.Subject, "1")
	assert.Equal(t, token.IssuedAt, currentTime)
	assert.Equal(t, token.ExpirationTime, expirationTime)
	assert.Equal(t, token.DeviceUUID, "1")
	assert.Equal(t, token.DeviceUserAgent, "1")
}

func Test_NewAccessTokenByJWT(t *testing.T) {
	config := configs.NewBaseConfig()

	currentTime := time.Now().Truncate(time.Second)

	token, err := NewAccessToken(
		config,
		"1",
		"1",
		"1",
		currentTime,
	)

	assert.Nil(t, err)
	assert.NotNil(t, token)

	tokenToJWT, err := token.GetJWT()

	assert.NoError(t, err)
	assert.NotNil(t, tokenToJWT)

	tokenByJWT, err := NewAccessTokenByJWT(config, tokenToJWT)

	assert.NoError(t, err)
	assert.NotNil(t, tokenByJWT)

	assert.Equal(t, tokenByJWT.Issuer, token.Issuer)
	assert.Equal(t, tokenByJWT.Audience, token.Audience)
	assert.Equal(t, tokenByJWT.Subject, token.Subject)
	assert.Equal(t, tokenByJWT.IssuedAt, token.IssuedAt)
	assert.Equal(t, tokenByJWT.ExpirationTime, token.ExpirationTime)
	assert.Equal(t, tokenByJWT.DeviceUUID, token.DeviceUUID)
	assert.Equal(t, tokenByJWT.DeviceUserAgent, token.DeviceUserAgent)

	token.IssuedAt = time.
		Now().
		Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes())).
		Truncate(time.Second)

	tokenToJWT, err = token.GetJWT()

	assert.NoError(t, err)
	assert.NotNil(t, tokenToJWT)

	tokenByJWT, err = NewAccessTokenByJWT(config, tokenToJWT)

	assert.Error(t, err)
	assert.Nil(t, tokenByJWT)

	token.ExpirationTime = time.
		Now().
		Add(-(time.Minute * time.Duration(config.GetExpirationAccessInMinutes()))).
		Truncate(time.Second)

	tokenToJWT, err = token.GetJWT()

	assert.NoError(t, err)
	assert.NotNil(t, tokenToJWT)

	tokenByJWT, err = NewAccessTokenByJWT(config, tokenToJWT)

	assert.Error(t, err)
	assert.Nil(t, tokenByJWT)

	tokenByJWT, err = NewAccessTokenByJWT(config, "0")

	assert.Error(t, err)
	assert.Nil(t, tokenByJWT)

	tokenByJWT, err = NewAccessTokenByJWT(config, "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.0-0-0-0-0-0")

	assert.Error(t, err)
	assert.Nil(t, tokenByJWT)
}
