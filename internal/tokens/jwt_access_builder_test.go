package tokens

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewJwtAccessTokenBuilder(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)
}

func TestJwtAccessTokenBuilder_New(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilderOne := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilderOne)

	jwtAccessTokenBuilderOne.SetSubject("1")

	jwtAccessTokenBuilderTwo := jwtAccessTokenBuilderOne.New()

	assert.NotNil(t, jwtAccessTokenBuilderTwo)
	assert.NotEqual(t, jwtAccessTokenBuilderOne, jwtAccessTokenBuilderTwo)
}

func TestJwtAccessTokenBuilder_NewFromJwtString(t *testing.T) {
	config := configs.NewBaseConfig()

	timeNow := time.Now().Truncate(time.Second)

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config).New()

	expirationTime := timeNow.
		Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes()))

	jwtAccessToken, err := jwtAccessTokenBuilder.
		New().
		SetSubject("1").
		SetDeviceUUID("1").
		SetDeviceUserAgent("1").
		SetIssuedAt(int(timeNow.Unix())).
		SetExpirationTime(int(expirationTime.Unix())).
		Build()

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtAccessToken)

	tokenToString, err := jwtAccessToken.ToString()

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenToString)

	jwtAccessTokenBuilder, err = jwtAccessTokenBuilder.NewFromJwtString(tokenToString)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtAccessTokenBuilder)

	token, err := jwtAccessTokenBuilder.Build()

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	assert.Equal(t, jwtAccessToken.GetIssuer(), token.GetIssuer())
	assert.Equal(t, jwtAccessToken.GetAudience(), token.GetAudience())
	assert.Equal(t, jwtAccessToken.GetSubject(), token.GetSubject())
	assert.Equal(t, jwtAccessToken.GetIssuedAt(), token.GetIssuedAt())
	assert.Equal(t, jwtAccessToken.GetExpirationTime(), token.GetExpirationTime())
	assert.Equal(t, jwtAccessToken.GetDeviceUUID(), token.GetDeviceUUID())
	assert.Equal(t, jwtAccessToken.GetDeviceUserAgent(), token.GetDeviceUserAgent())

	jwtAccessTokenBuilderTemp := jwtAccessTokenBuilder

	timeNow = time.
		Now().
		Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes())).
		Truncate(time.Second)

	jwtAccessToken, err = jwtAccessTokenBuilderTemp.
		SetIssuedAt(int(timeNow.Unix())).
		Build()

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtAccessToken)

	tokenToString, err = jwtAccessToken.ToString()

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenToString)

	jwtAccessTokenBuilderTemp, err = jwtAccessTokenBuilderTemp.NewFromJwtString(tokenToString)

	assert.Error(t, err)
	assert.Empty(t, jwtAccessTokenBuilderTemp)

	jwtAccessTokenBuilderTemp = jwtAccessTokenBuilder

	timeNow = time.
		Now().
		Add(-(time.Minute * time.Duration(config.GetExpirationAccessInMinutes()))).
		Truncate(time.Second)

	jwtAccessToken, err = jwtAccessTokenBuilderTemp.
		SetExpirationTime(int(timeNow.Unix())).
		Build()

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtAccessToken)

	tokenToString, err = jwtAccessToken.ToString()

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenToString)

	jwtAccessTokenBuilderTemp, err = jwtAccessTokenBuilder.NewFromJwtString(tokenToString)

	assert.Error(t, err)
	assert.Empty(t, jwtAccessTokenBuilderTemp)

	jwtAccessTokenBuilderTemp, err = jwtAccessTokenBuilder.NewFromJwtString("0")

	assert.Error(t, err)
	assert.Empty(t, jwtAccessTokenBuilderTemp)

	jwtAccessTokenBuilderTemp, err = jwtAccessTokenBuilder.NewFromJwtString("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.0-0-0-0-0-0")

	assert.Error(t, err)
	assert.Empty(t, jwtAccessTokenBuilderTemp)
}

func TestJwtAccessTokenBuilder_Build(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetSubject("1")
	jwtAccessTokenBuilder.SetIssuedAt(1)
	jwtAccessTokenBuilder.SetExpirationTime(1)
	jwtAccessTokenBuilder.SetDeviceUUID("1")
	jwtAccessTokenBuilder.SetDeviceUserAgent("1")

	jwtAccessToken, err := jwtAccessTokenBuilder.BuildToJwt()

	assert.NoError(t, err)
	assert.NotNil(t, jwtAccessToken)
}

func TestJwtAccessTokenBuilder_BuildToJwt(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessToken, err := jwtAccessTokenBuilder.BuildToJwt()

	assert.Error(t, err)
	assert.Nil(t, jwtAccessToken)

	jwtAccessTokenBuilder.SetSubject("1")

	jwtAccessToken, err = jwtAccessTokenBuilder.BuildToJwt()

	assert.Error(t, err)
	assert.Nil(t, jwtAccessToken)

	jwtAccessTokenBuilder.SetIssuedAt(1)

	jwtAccessToken, err = jwtAccessTokenBuilder.BuildToJwt()

	assert.Error(t, err)
	assert.Nil(t, jwtAccessToken)

	jwtAccessTokenBuilder.SetExpirationTime(1)

	jwtAccessToken, err = jwtAccessTokenBuilder.BuildToJwt()

	assert.Error(t, err)
	assert.Nil(t, jwtAccessToken)

	jwtAccessTokenBuilder.SetDeviceUUID("1")

	jwtAccessToken, err = jwtAccessTokenBuilder.BuildToJwt()

	assert.Error(t, err)
	assert.Nil(t, jwtAccessToken)

	jwtAccessTokenBuilder.SetDeviceUserAgent("1")

	jwtAccessToken, err = jwtAccessTokenBuilder.BuildToJwt()

	assert.NoError(t, err)
	assert.NotNil(t, jwtAccessToken)
}

func TestJwtAccessTokenBuilder_SetIssuer(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetIssuer("1")

	assert.Equal(t, "1", jwtAccessTokenBuilder.accessToken.GetIssuer())

	jwtAccessTokenBuilder.SetIssuer("2")

	assert.Equal(t, "2", jwtAccessTokenBuilder.accessToken.GetIssuer())
}

func TestJwtAccessTokenBuilder_SetAudience(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetAudience("1")

	assert.Equal(t, "1", jwtAccessTokenBuilder.accessToken.GetAudience())

	jwtAccessTokenBuilder.SetAudience("2")

	assert.Equal(t, "2", jwtAccessTokenBuilder.accessToken.GetAudience())
}

func TestJwtAccessTokenBuilder_SetSubject(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetSubject("1")

	assert.Equal(t, "1", jwtAccessTokenBuilder.accessToken.GetSubject())

	jwtAccessTokenBuilder.SetSubject("2")

	assert.Equal(t, "2", jwtAccessTokenBuilder.accessToken.GetSubject())
}

func TestJwtAccessTokenBuilder_SetDeviceUUID(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetDeviceUUID("1")

	assert.Equal(t, "1", jwtAccessTokenBuilder.accessToken.GetDeviceUUID())

	jwtAccessTokenBuilder.SetDeviceUUID("2")

	assert.Equal(t, "2", jwtAccessTokenBuilder.accessToken.GetDeviceUUID())
}

func TestJwtAccessTokenBuilder_SetDeviceUserAgent(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetDeviceUserAgent("1")

	assert.Equal(t, "1", jwtAccessTokenBuilder.accessToken.GetDeviceUserAgent())

	jwtAccessTokenBuilder.SetDeviceUserAgent("2")

	assert.Equal(t, "2", jwtAccessTokenBuilder.accessToken.GetDeviceUserAgent())
}

func TestJwtAccessTokenBuilder_SetIssuedAt(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetIssuedAt(1)

	assert.Equal(t, 1, jwtAccessTokenBuilder.accessToken.GetIssuedAt())

	jwtAccessTokenBuilder.SetIssuedAt(2)

	assert.Equal(t, 2, jwtAccessTokenBuilder.accessToken.GetIssuedAt())
}

func TestJwtAccessTokenBuilder_SetExpirationTime(t *testing.T) {
	config := configs.NewBaseConfig()

	jwtAccessTokenBuilder := NewJwtAccessTokenBuilder(config)

	assert.NotNil(t, jwtAccessTokenBuilder)

	jwtAccessTokenBuilder.SetExpirationTime(1)

	assert.Equal(t, 1, jwtAccessTokenBuilder.accessToken.GetExpirationTime())

	jwtAccessTokenBuilder.SetExpirationTime(2)

	assert.Equal(t, 2, jwtAccessTokenBuilder.accessToken.GetExpirationTime())
}
