package requests

import (
	"github.com/fromsi/jwt-oauth-sso/internal/validator_rules"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewPasswordResetWithTokenRequestBody(t *testing.T) {
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("password", validator_rules.Password)

		assert.NoError(t, err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest("POST", "", strings.NewReader(tt.body))
			c.Request.Header.Set("Content-Type", "application/json")

			requestBody, err := NewPasswordResetWithTokenRequestBody(c)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, requestBody)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, requestBody)
				assert.NotEmpty(t, requestBody.Token)
				assert.NotEmpty(t, requestBody.NewPassword)
			}
		})
	}
}
