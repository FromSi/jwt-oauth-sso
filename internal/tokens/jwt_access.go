package tokens

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/golang-jwt/jwt/v5"
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

type JwtAccessToken struct {
	config          configs.TokenConfig
	issuer          string
	audience        string
	subject         string
	deviceUUID      string
	deviceUserAgent string
	issuedAt        int
	expirationTime  int
}

func (receiver *JwtAccessToken) ToString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CommonJWTClaimIssuer:          receiver.issuer,
		CommonJWTClaimAudience:        receiver.audience,
		CommonJWTClaimSubject:         receiver.subject,
		CommonJWTClaimIssuedAt:        receiver.issuedAt,
		CommonJWTClaimExpirationTime:  receiver.expirationTime,
		CommonJWTClaimDeviceUUID:      receiver.deviceUUID,
		CommonJWTClaimDeviceUserAgent: receiver.deviceUserAgent,
	})

	return token.SignedString([]byte(receiver.config.GetSecretKey()))
}

func (receiver *JwtAccessToken) GetIssuer() string {
	return receiver.issuer
}

func (receiver *JwtAccessToken) GetAudience() string {
	return receiver.audience
}

func (receiver *JwtAccessToken) GetSubject() string {
	return receiver.subject
}

func (receiver *JwtAccessToken) GetDeviceUUID() string {
	return receiver.deviceUUID
}

func (receiver *JwtAccessToken) GetDeviceUserAgent() string {
	return receiver.deviceUserAgent
}

func (receiver *JwtAccessToken) GetIssuedAt() int {
	return receiver.issuedAt
}

func (receiver *JwtAccessToken) GetExpirationTime() int {
	return receiver.expirationTime
}

func (receiver *JwtAccessToken) SetIssuer(value string) {
	receiver.issuer = value
}

func (receiver *JwtAccessToken) SetAudience(value string) {
	receiver.audience = value
}

func (receiver *JwtAccessToken) SetSubject(value string) {
	receiver.subject = value
}

func (receiver *JwtAccessToken) SetDeviceUUID(value string) {
	receiver.deviceUUID = value
}

func (receiver *JwtAccessToken) SetDeviceUserAgent(value string) {
	receiver.deviceUserAgent = value
}

func (receiver *JwtAccessToken) SetIssuedAt(value int) {
	receiver.issuedAt = value
}

func (receiver *JwtAccessToken) SetExpirationTime(value int) {
	receiver.expirationTime = value
}
