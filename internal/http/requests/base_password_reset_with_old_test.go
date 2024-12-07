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

func Test_NewBasePasswordResetWithOldRequest_And_NewBasePasswordResetWithOldRequestBody_And_GetBody(t *testing.T) {
	requestBody := NewBasePasswordResetWithOldRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBasePasswordResetWithOldRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())
}

func TestNewBasePasswordResetWithOldRequestBody_GetOldPassword(t *testing.T) {
	requestBody := NewBasePasswordResetWithOldRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.OldPassword = "1"

	assert.Equal(t, "1", requestBody.GetOldPassword())

	requestBody.OldPassword = "2"

	assert.Equal(t, "2", requestBody.GetOldPassword())
}

func TestNewBasePasswordResetWithOldRequestBody_GetNewPassword(t *testing.T) {
	requestBody := NewBasePasswordResetWithOldRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.NewPassword = "1"

	assert.Equal(t, "1", requestBody.GetNewPassword())

	requestBody.NewPassword = "2"

	assert.Equal(t, "2", requestBody.GetNewPassword())
}

func TestNewBasePasswordResetWithOldRequest_Make(t *testing.T) {
	requestConstructorBody := NewBasePasswordResetWithOldRequestBody()

	assert.NotNil(t, requestConstructorBody)

	requestConstructor := NewBasePasswordResetWithOldRequest(requestConstructorBody)

	assert.NotNil(t, requestConstructor)

	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{"newPassword": "validPass123!", "oldPassword": "validPass123!"}`,
			error: false,
		},
		{
			name:  "Invalid newPassword",
			body:  `{"newPassword": "123", "oldPassword": "validPass123!"}`,
			error: true,
		},
		{
			name:  "Invalid oldPassword",
			body:  `{"newPassword": "validPass123!", "oldPassword": "123"}`,
			error: true,
		},
		{
			name:  "Missing newPassword",
			body:  `{"oldPassword": "validPass123!"}`,
			error: true,
		},
		{
			name:  "Missing oldPassword",
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

			request, errResponse := requestConstructor.Make(c)

			if tt.error {
				assert.NotEmpty(t, errResponse)
				assert.Empty(t, request)
			} else {
				assert.Empty(t, errResponse)
				assert.NotEmpty(t, request)
				assert.NotEmpty(t, request.GetBody())
				assert.NotEmpty(t, request.GetBody().GetOldPassword())
				assert.NotEmpty(t, request.GetBody().GetNewPassword())
			}
		})
	}
}
