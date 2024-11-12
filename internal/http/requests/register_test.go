package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/validator_rules"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewRegisterRequest_And_NewRegisterRequestBody(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		ip        string
		userAgent string
		error     bool
	}{
		{
			name:      "Valid request",
			body:      `{"email": "test@example.com", "password": "validPass123!"}`,
			ip:        "127.0.0.1",
			userAgent: "Mozilla/5.0",
			error:     false,
		},
		{
			name:      "Invalid email",
			body:      `{"email": "invalid-email", "password": "validPass123!"}`,
			ip:        "127.0.0.1",
			userAgent: "Mozilla/5.0",
			error:     true,
		},
		{
			name:      "Invalid password",
			body:      `{"email": "test@example.com", "password": "123"}`,
			ip:        "127.0.0.1",
			userAgent: "Mozilla/5.0",
			error:     true,
		},
		{
			name:      "Missing email",
			body:      `{"password": "validPass123!"}`,
			ip:        "127.0.0.1",
			userAgent: "Mozilla/5.0",
			error:     true,
		},
		{
			name:      "Missing password",
			body:      `{"email": "test@example.com"}`,
			ip:        "127.0.0.1",
			userAgent: "Mozilla/5.0",
			error:     true,
		},
		{
			name:      "Empty request",
			body:      `{}`,
			ip:        "127.0.0.1",
			userAgent: "Mozilla/5.0",
			error:     true,
		},
	}

	gin.SetMode(gin.TestMode)

	err := validator_rules.BindPassword()

	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest("POST", "", strings.NewReader(tt.body))
			c.Request.Header.Set("Content-Type", "application/json")
			c.Request.Header.Set("User-Agent", tt.userAgent)
			c.Request.Header.Set("X-Real-Ip", tt.ip)

			request, errResponse := NewRegisterRequest(c)

			if tt.error {
				assert.NotEmpty(t, errResponse)
				assert.Empty(t, request)
			} else {
				assert.Empty(t, errResponse)
				assert.NotEmpty(t, request)
				assert.NotEmpty(t, request.Body)
				assert.NotEmpty(t, request.Body.Email)
				assert.NotEmpty(t, request.Body.Password)
				assert.Equal(t, tt.ip, request.IP)
				assert.Equal(t, tt.userAgent, request.UserAgent)
			}
		})
	}
}
