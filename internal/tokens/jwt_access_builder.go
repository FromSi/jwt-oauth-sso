package tokens

import (
	"errors"
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtAccessTokenBuilder struct {
	accessToken JwtAccessToken
}

func NewJwtAccessTokenBuilder(config configs.TokenConfig) *JwtAccessTokenBuilder {
	return &JwtAccessTokenBuilder{
		accessToken: JwtAccessToken{config: config},
	}
}

func (receiver *JwtAccessTokenBuilder) New() AccessTokenBuilder {
	jwtAccessTokenBuilder := &JwtAccessTokenBuilder{
		accessToken: JwtAccessToken{config: receiver.accessToken.config},
	}

	jwtAccessTokenBuilder.
		accessToken.
		SetIssuer(receiver.accessToken.config.GetIssuerName())

	jwtAccessTokenBuilder.
		accessToken.
		SetAudience(receiver.accessToken.config.GetAudienceName())

	return &JwtAccessTokenBuilder{
		accessToken: JwtAccessToken{config: receiver.accessToken.config},
	}
}

func (receiver *JwtAccessTokenBuilder) NewFromJwtString(token string) (AccessTokenBuilder, error) {
	tokenJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"unexpected signing method: %v",
				token.Header["alg"],
			)
		}
		return []byte(receiver.accessToken.config.GetSecretKey()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}

		return nil, errors.New("invalid token")
	}

	claims := tokenJwt.Claims.(jwt.MapClaims)

	issuedAt := time.
		Unix(int64(claims[CommonJWTClaimIssuedAt].(float64)), 0).
		Unix()

	expirationTime := time.
		Unix(int64(claims[CommonJWTClaimExpirationTime].(float64)), 0).
		Unix()

	accessToken := JwtAccessToken{
		config:          receiver.accessToken.config,
		issuer:          claims[CommonJWTClaimIssuer].(string),
		audience:        claims[CommonJWTClaimAudience].(string),
		subject:         claims[CommonJWTClaimSubject].(string),
		issuedAt:        int(issuedAt),
		expirationTime:  int(expirationTime),
		deviceUUID:      claims[CommonJWTClaimDeviceUUID].(string),
		deviceUserAgent: claims[CommonJWTClaimDeviceUserAgent].(string),
	}

	isIssued := time.Now().Before(time.Unix(int64(accessToken.GetIssuedAt()), 0))

	if isIssued {
		return nil, errors.New("token used before issued")
	}

	return &JwtAccessTokenBuilder{
		accessToken: accessToken,
	}, nil
}

func (receiver *JwtAccessTokenBuilder) Build() (AccessToken, error) {
	return receiver.BuildToJwt()
}

func (receiver *JwtAccessTokenBuilder) BuildToJwt() (*JwtAccessToken, error) {
	if len(receiver.accessToken.GetSubject()) == 0 {
		return nil, errors.New("subject must not be empty")
	}

	if receiver.accessToken.GetIssuedAt() == 0 {
		return nil, errors.New("issuedAt must not be empty")
	}

	if receiver.accessToken.GetExpirationTime() == 0 {
		return nil, errors.New("expirationTime must not be empty")
	}

	if len(receiver.accessToken.GetDeviceUUID()) == 0 {
		return nil, errors.New("deviceUUID must not be empty")
	}

	if len(receiver.accessToken.GetDeviceUserAgent()) == 0 {
		return nil, errors.New("deviceUserAgent must not be empty")
	}

	return &receiver.accessToken, nil
}

func (receiver *JwtAccessTokenBuilder) SetIssuer(value string) AccessTokenBuilder {
	receiver.accessToken.SetIssuer(value)

	return receiver
}

func (receiver *JwtAccessTokenBuilder) SetAudience(value string) AccessTokenBuilder {
	receiver.accessToken.SetAudience(value)

	return receiver
}

func (receiver *JwtAccessTokenBuilder) SetSubject(value string) AccessTokenBuilder {
	receiver.accessToken.SetSubject(value)

	return receiver
}

func (receiver *JwtAccessTokenBuilder) SetDeviceUUID(value string) AccessTokenBuilder {
	receiver.accessToken.SetDeviceUUID(value)

	return receiver
}

func (receiver *JwtAccessTokenBuilder) SetDeviceUserAgent(value string) AccessTokenBuilder {
	receiver.accessToken.SetDeviceUserAgent(value)

	return receiver
}

func (receiver *JwtAccessTokenBuilder) SetIssuedAt(value int) AccessTokenBuilder {
	receiver.accessToken.SetIssuedAt(value)

	return receiver
}

func (receiver *JwtAccessTokenBuilder) SetExpirationTime(value int) AccessTokenBuilder {
	receiver.accessToken.SetExpirationTime(value)

	return receiver
}
