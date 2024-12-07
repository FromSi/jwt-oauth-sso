package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewBaseLogoutAllRequest_And_NewBaseLogoutAllRequestBody_And_GetBody(t *testing.T) {
	requestBody := NewBaseLogoutAllRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseLogoutAllRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())
}

func TestNewBaseLogoutAllRequest_Make(t *testing.T) {
	requestConstructorBody := NewBaseLogoutAllRequestBody()

	assert.NotNil(t, requestConstructorBody)

	requestConstructor := NewBaseLogoutAllRequest(requestConstructorBody)

	assert.NotNil(t, requestConstructor)

	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{}`,
			error: false,
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest("POST", "", strings.NewReader(tt.body))
			c.Request.Header.Set("Content-Type", "application/json")

			request := requestConstructor.Make(c)

			if tt.error {
				assert.Empty(t, request)
			} else {
				assert.NotNil(t, request)
				assert.NotNil(t, request.GetBody())
			}
		})
	}
}
