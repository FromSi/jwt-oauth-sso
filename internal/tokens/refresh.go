package tokens

import (
	"errors"
	"fmt"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	CommonJWTClaimIssuer         = "iss"
	CommonJWTClaimAudience       = "aud"
	CommonJWTClaimSubject        = "sub"
	CommonJWTClaimIssuedAt       = "iat"
	CommonJWTClaimExpirationTime = "exp"
	CommonJWTClaimDeviceUUID     = "device_uuid"
	CommonJWTClaimDeviceAgent    = "device_agent"
)

type RefreshToken struct {
	Issuer         string
	Audience       string
	Subject        string
	IssuedAt       time.Time
	ExpirationTime time.Time
	DeviceUUID     string
	DeviceAgent    string
	secretKey      string
}

func NewRefreshToken(
	config configs.TokenConfig,
	subject string,
	deviceUUID string,
	deviceAgent string,
	currentTime time.Time,
) (*RefreshToken, error) {
	return &RefreshToken{
		Issuer:         config.GetIssuerName(),
		Audience:       config.GetAudienceName(),
		Subject:        subject,
		IssuedAt:       currentTime,
		ExpirationTime: currentTime.Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes())),
		DeviceUUID:     deviceUUID,
		DeviceAgent:    deviceAgent,
		secretKey:      config.GetSecretKey(),
	}, nil
}

func NewRefreshTokenByJWT(config configs.TokenConfig, tokenJWT string) (*RefreshToken, error) {
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

	refreshToken := RefreshToken{
		Issuer:         claims[CommonJWTClaimIssuer].(string),
		Audience:       claims[CommonJWTClaimAudience].(string),
		Subject:        claims[CommonJWTClaimSubject].(string),
		IssuedAt:       time.Unix(int64(claims[CommonJWTClaimIssuedAt].(float64)), 0),
		ExpirationTime: time.Unix(int64(claims[CommonJWTClaimExpirationTime].(float64)), 0),
		DeviceUUID:     claims[CommonJWTClaimDeviceUUID].(string),
		DeviceAgent:    claims[CommonJWTClaimDeviceAgent].(string),
		secretKey:      config.GetSecretKey(),
	}

	if time.Now().Before(refreshToken.IssuedAt) {
		return nil, errors.New("token used before issued")
	}

	if time.Now().After(refreshToken.ExpirationTime) {
		return nil, errors.New("token has expired")
	}

	return &refreshToken, nil
}

func (receiver RefreshToken) GetJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CommonJWTClaimIssuer:         receiver.Issuer,
		CommonJWTClaimAudience:       receiver.Audience,
		CommonJWTClaimSubject:        receiver.Subject,
		CommonJWTClaimIssuedAt:       receiver.IssuedAt.Unix(),
		CommonJWTClaimExpirationTime: receiver.ExpirationTime.Unix(),
		CommonJWTClaimDeviceUUID:     receiver.DeviceUUID,
		CommonJWTClaimDeviceAgent:    receiver.DeviceAgent,
	})

	return token.SignedString([]byte(receiver.secretKey))
}