package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_NewBearerAuthRequestHeader(t *testing.T) {
	config := configs.NewBaseConfig()

	validToken, _ := tokens.NewAccessToken(
		config,
		"1",
		"1",
		"1",
		time.Now(),
	)

	validTokenToJWT, _ := validToken.GetJWT()

	invalidToken, _ := tokens.NewAccessToken(
		config,
		"1",
		"1",
		"1",
		time.Unix(1, 1),
	)

	invalidTokenToJWT, _ := invalidToken.GetJWT()

	tests := []struct {
		name        string
		accessToken string
		error       bool
	}{
		{
			name:        "Valid request",
			accessToken: validTokenToJWT,
			error:       false,
		},
		{
			name:        "Invalid request",
			accessToken: invalidTokenToJWT,
			error:       true,
		},
		{
			name:        "Empty request",
			accessToken: "",
			error:       true,
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest("POST", "", strings.NewReader(""))

			if len(tt.accessToken) != 0 {
				c.Request.Header.Set("Authorization", "Bearer "+tt.accessToken)
			}

			request, err := NewBearerAuthRequestHeader(c, config)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, request)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, request)
			}
		})
	}
}
