package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewLogoutDeviceRequestBody(t *testing.T) {
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

			requestBody, errResponse := NewLogoutDeviceRequestBody(c)

			if tt.error {
				assert.NotNil(t, errResponse)
				assert.Nil(t, requestBody)
			} else {
				assert.Nil(t, errResponse)
				assert.NotNil(t, requestBody)
				assert.NotEmpty(t, requestBody.DeviceUUID)
			}
		})
	}
}
