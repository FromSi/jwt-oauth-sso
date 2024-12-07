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

func Test_NewBaseLoginRequest_And_NewBaseLoginRequestBody_And_GetBody(t *testing.T) {
	requestBody := NewBaseLoginRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseLoginRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())
}

func TestNewBaseLogoutDeviceRequestBody_GetEmail(t *testing.T) {
	requestBody := NewBaseLoginRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.Email = "1"

	assert.Equal(t, "1", requestBody.GetEmail())

	requestBody.Email = "2"

	assert.Equal(t, "2", requestBody.GetEmail())
}

func TestNewBaseLogoutDeviceRequestBody_GetPassword(t *testing.T) {
	requestBody := NewBaseLoginRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.Password = "1"

	assert.Equal(t, "1", requestBody.GetPassword())

	requestBody.Password = "2"

	assert.Equal(t, "2", requestBody.GetPassword())
}

func TestNewBaseLogoutDeviceRequest_GetIP(t *testing.T) {
	requestBody := NewBaseLoginRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseLoginRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())

	request.IP = "1"

	assert.Equal(t, "1", request.GetIP())

	request.IP = "2"

	assert.Equal(t, "2", request.GetIP())
}

func TestNewBaseLogoutDeviceRequest_GetUserAgent(t *testing.T) {
	requestBody := NewBaseLoginRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseLoginRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())

	request.UserAgent = "1"

	assert.Equal(t, "1", request.GetUserAgent())

	request.UserAgent = "2"

	assert.Equal(t, "2", request.GetUserAgent())
}

func TestNewBaseLoginRequest_Make(t *testing.T) {
	requestConstructorBody := NewBaseLoginRequestBody()

	assert.NotNil(t, requestConstructorBody)

	requestConstructor := NewBaseLoginRequest(requestConstructorBody)

	assert.NotNil(t, requestConstructor)

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
			c.Request.Header.Set("X-Real-Ip", tt.ip)
			c.Request.Header.Set("User-Agent", tt.userAgent)

			request, errResponse := requestConstructor.Make(c)

			if tt.error {
				assert.NotEmpty(t, errResponse)
				assert.Empty(t, request)
			} else {
				assert.Empty(t, errResponse)
				assert.NotEmpty(t, request)
				assert.NotEmpty(t, request.GetBody())
				assert.NotEmpty(t, request.GetBody().GetEmail())
				assert.NotEmpty(t, request.GetBody().GetPassword())
				assert.Equal(t, tt.ip, request.GetIP())
				assert.Equal(t, tt.userAgent, request.GetUserAgent())
			}
		})
	}
}
