package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewBaseSendResetTokenRequest_And_NewBaseSendResetTokenRequestBody_And_GetBody(t *testing.T) {
	requestBody := NewBaseSendResetTokenRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseSendResetTokenRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())
}

func TestNewBaseSendResetTokenRequestBody_GetEmail(t *testing.T) {
	requestBody := NewBaseSendResetTokenRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.Email = "1"

	assert.Equal(t, "1", requestBody.GetEmail())

	requestBody.Email = "2"

	assert.Equal(t, "2", requestBody.GetEmail())
}

func TestNewBaseSendResetTokenRequest_Make(t *testing.T) {
	requestConstructorBody := NewBaseSendResetTokenRequestBody()

	assert.NotNil(t, requestConstructorBody)

	requestConstructor := NewBaseSendResetTokenRequest(requestConstructorBody)

	assert.NotNil(t, requestConstructor)

	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{"email": "test@fromsi.net"}`,
			error: false,
		},
		{
			name:  "Invalid email",
			body:  `{"email": "test"}`,
			error: true,
		},
		{
			name:  "Empty request",
			body:  `{}`,
			error: true,
		},
	}

	gin.SetMode(gin.TestMode)

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
				assert.NotEmpty(t, request.GetBody().GetEmail())
			}
		})
	}
}
