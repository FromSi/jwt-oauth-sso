package tokens

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewAccessToken(t *testing.T) {
	tests := []struct {
		name            string
		subject         string
		deviceUUID      string
		deviceUserAgent string
		currentTime     time.Time
		error           bool
	}{
		{
			name:            "Valid data",
			subject:         "1",
			deviceUUID:      "1",
			deviceUserAgent: "1",
			currentTime:     time.Now().Truncate(time.Second),
			error:           false,
		},
	}

	config := configs.NewBaseConfig()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := NewAccessToken(config, tt.subject, tt.deviceUUID, tt.deviceUserAgent, tt.currentTime)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, token)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, token.Issuer, config.GetIssuerName())
				assert.Equal(t, token.Audience, config.GetAudienceName())
				assert.Equal(t, token.Subject, tt.subject)
				assert.Equal(t, token.IssuedAt, tt.currentTime)
				assert.Equal(t, token.ExpirationTime, tt.currentTime.Add(time.Minute*time.Duration(config.GetExpirationAccessInMinutes())))
				assert.Equal(t, token.DeviceUUID, tt.deviceUUID)
				assert.Equal(t, token.DeviceUserAgent, tt.deviceUserAgent)
			}
		})
	}
}

func Test_NewAccessTokenByJWT(t *testing.T) {
	config := configs.NewBaseConfig()

	tests := []struct {
		name        string
		subject     string
		deviceUUID  string
		deviceAgent string
		currentTime time.Time
		error       bool
	}{
		{
			name:        "Valid data",
			subject:     "1",
			deviceUUID:  "1",
			deviceAgent: "1",
			currentTime: time.Now().Truncate(time.Second),
			error:       false,
		},
		{
			name:        "Invalid token used before issued",
			subject:     "2",
			deviceUUID:  "2",
			deviceAgent: "2",
			currentTime: time.Now().Add(time.Minute * time.Duration(config.GetExpirationAccessInMinutes())).Truncate(time.Second),
			error:       true,
		},
		{
			name:        "Invalid token is expired",
			subject:     "3",
			deviceUUID:  "3",
			deviceAgent: "3",
			currentTime: time.Now().Add(-(time.Minute * time.Duration(config.GetExpirationAccessInMinutes()))).Truncate(time.Second),
			error:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := NewAccessToken(config, tt.subject, tt.deviceUUID, tt.deviceAgent, tt.currentTime)

			assert.NoError(t, err)

			tokenToJWT, err := token.GetJWT()

			assert.NoError(t, err)
			assert.NotNil(t, tokenToJWT)

			tokenByJWT, err := NewAccessTokenByJWT(config, tokenToJWT)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, tokenByJWT)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tokenByJWT.Issuer, token.Issuer)
				assert.Equal(t, tokenByJWT.Audience, token.Audience)
				assert.Equal(t, tokenByJWT.Subject, token.Subject)
				assert.Equal(t, tokenByJWT.IssuedAt, token.IssuedAt)
				assert.Equal(t, tokenByJWT.ExpirationTime, token.ExpirationTime)
				assert.Equal(t, tokenByJWT.DeviceUUID, token.DeviceUUID)
				assert.Equal(t, tokenByJWT.DeviceUserAgent, token.DeviceUserAgent)
			}
		})
	}
}
