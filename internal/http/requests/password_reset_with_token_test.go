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

func Test_NewPasswordResetWithTokenRequest_And_NewPasswordResetWithTokenRequestBody(t *testing.T) {
	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{"token": "2e79c328-d6a9-45d4-92e7-2677aa36f6c3", "newPassword": "validPass123!"}`,
			error: false,
		},
		{
			name:  "Invalid token",
			body:  `{"token": "2e79c328-45d4-92e7-2677aa36f6c3", "newPassword": "validPass123!"}`,
			error: true,
		},
		{
			name:  "Invalid newPassword",
			body:  `{"token": "2e79c328-d6a9-45d4-92e7-2677aa36f6c3", "newPassword": "123"}`,
			error: true,
		},
		{
			name:  "Missing newPassword",
			body:  `{"token": "2e79c328-d6a9-45d4-92e7-2677aa36f6c3"}`,
			error: true,
		},
		{
			name:  "Missing token",
			body:  `{"newPassword": "validPass123!"}`,
			error: true,
		},
		{
			name:  "Empty request",
			body:  `{}`,
			error: true,
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

			request, errResponse := NewPasswordResetWithTokenRequest(c)

			if tt.error {
				assert.NotEmpty(t, errResponse)
				assert.Empty(t, request)
			} else {
				assert.Empty(t, errResponse)
				assert.NotEmpty(t, request)
				assert.NotEmpty(t, request.Body)
				assert.NotEmpty(t, request.Body.Token)
				assert.NotEmpty(t, request.Body.NewPassword)
			}
		})
	}
}
