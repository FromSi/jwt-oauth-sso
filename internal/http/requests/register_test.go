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

func Test_NewRegisterRequestBody(t *testing.T) {
	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{"email": "test@example.com", "password": "validPass123!"}`,
			error: false,
		},
		{
			name:  "Invalid email",
			body:  `{"email": "invalid-email", "password": "validPass123!"}`,
			error: true,
		},
		{
			name:  "Invalid password",
			body:  `{"email": "test@example.com", "password": "123"}`,
			error: true,
		},
		{
			name:  "Missing email",
			body:  `{"password": "validPass123!"}`,
			error: true,
		},
		{
			name:  "Missing password",
			body:  `{"email": "test@example.com"}`,
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

			requestBody, err := NewRegisterRequestBody(c)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, requestBody)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, requestBody)
				assert.NotEmpty(t, requestBody.Email)
				assert.NotEmpty(t, requestBody.Password)
			}
		})
	}
}
