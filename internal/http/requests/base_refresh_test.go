package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewBaseRefreshRequest_And_NewBaseRefreshRequestBody_And_GetBody(t *testing.T) {
	requestBody := NewBaseRefreshRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseRefreshRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())
}

func TestNewBaseRefreshRequestBody_GetRefreshToken(t *testing.T) {
	requestBody := NewBaseRefreshRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.RefreshToken = "1"

	assert.Equal(t, "1", requestBody.GetRefreshToken())

	requestBody.RefreshToken = "2"

	assert.Equal(t, "2", requestBody.GetRefreshToken())
}

func TestNewBaseRefreshRequest_Make(t *testing.T) {
	requestConstructorBody := NewBaseRefreshRequestBody()

	assert.NotNil(t, requestConstructorBody)

	requestConstructor := NewBaseRefreshRequest(requestConstructorBody)

	assert.NotNil(t, requestConstructor)

	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{"refreshToken": "09d0ce56-33ab-4fce-90bb-d2b4d6d844ba"}`,
			error: false,
		},
		{
			name:  "Invalid refreshToken",
			body:  `{"refreshToken": "09d0ce56-33ab-4fce-90bb"}`,
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
			}
		})
	}
}
