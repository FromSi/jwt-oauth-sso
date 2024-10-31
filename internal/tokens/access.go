package tokens

import (
	"errors"
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	CommonJWTClaimIssuer          = "iss"
	CommonJWTClaimAudience        = "aud"
	CommonJWTClaimSubject         = "sub"
	CommonJWTClaimIssuedAt        = "iat"
	CommonJWTClaimExpirationTime  = "exp"
	CommonJWTClaimDeviceUUID      = "deviceUUID"
	CommonJWTClaimDeviceUserAgent = "deviceUserAgent"
)

type AccessToken struct {
	Issuer          string
	Audience        string
	Subject         string
	IssuedAt        time.Time
	ExpirationTime  time.Time
	DeviceUUID      string
	DeviceUserAgent string
	secretKey       string
}

func NewAccessToken(
	config configs.TokenConfig,
	subject string,
	deviceUUID string,
	deviceAgent string,
	currentTime time.Time,
) (*AccessToken, error) {
	return &AccessToken{
		Issuer:          config.GetIssuerName(),
		Audience:        config.GetAudienceName(),
		Subject:         subject,
		IssuedAt:        currentTime,
		ExpirationTime:  currentTime.Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes())),
		DeviceUUID:      deviceUUID,
		DeviceUserAgent: deviceAgent,
		secretKey:       config.GetSecretKey(),
	}, nil
}

func NewAccessTokenByJWT(config configs.TokenConfig, tokenJWT string) (*AccessToken, error) {
	token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetSecretKey()), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	accessToken := AccessToken{
		Issuer:          claims[CommonJWTClaimIssuer].(string),
		Audience:        claims[CommonJWTClaimAudience].(string),
		Subject:         claims[CommonJWTClaimSubject].(string),
		IssuedAt:        time.Unix(int64(claims[CommonJWTClaimIssuedAt].(float64)), 0),
		ExpirationTime:  time.Unix(int64(claims[CommonJWTClaimExpirationTime].(float64)), 0),
		DeviceUUID:      claims[CommonJWTClaimDeviceUUID].(string),
		DeviceUserAgent: claims[CommonJWTClaimDeviceUserAgent].(string),
		secretKey:       config.GetSecretKey(),
	}

	if time.Now().Before(accessToken.IssuedAt) {
		return nil, errors.New("token used before issued")
	}

	if time.Now().After(accessToken.ExpirationTime) {
		return nil, errors.New("token has expired")
	}

	return &accessToken, nil
}

func (receiver *AccessToken) GetJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CommonJWTClaimIssuer:          receiver.Issuer,
		CommonJWTClaimAudience:        receiver.Audience,
		CommonJWTClaimSubject:         receiver.Subject,
		CommonJWTClaimIssuedAt:        receiver.IssuedAt.Unix(),
		CommonJWTClaimExpirationTime:  receiver.ExpirationTime.Unix(),
		CommonJWTClaimDeviceUUID:      receiver.DeviceUUID,
		CommonJWTClaimDeviceUserAgent: receiver.DeviceUserAgent,
	})

	return token.SignedString([]byte(receiver.secretKey))
}
