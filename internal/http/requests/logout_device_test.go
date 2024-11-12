package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewLogoutDeviceRequest_And_NewLogoutDeviceRequestBody(t *testing.T) {
	tests := []struct {
		name  string
		body  string
		error bool
	}{
		{
			name:  "Valid request",
			body:  `{"deviceUuid": "2e79c328-d6a9-45d4-92e7-2677aa36f6c3"}`,
			error: false,
		},
		{
			name:  "Invalid deviceUuid",
			body:  `{"deviceUuid": "2e79c328-45d4-92e7-2677aa36f6c3"}`,
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

			request, errResponse := NewLogoutDeviceRequest(c)

			if tt.error {
				assert.NotEmpty(t, errResponse)
				assert.Empty(t, request)
			} else {
				assert.Empty(t, errResponse)
				assert.NotEmpty(t, request)
				assert.NotEmpty(t, request.Body)
				assert.NotEmpty(t, request.Body.DeviceUUID)
			}
		})
	}
}
