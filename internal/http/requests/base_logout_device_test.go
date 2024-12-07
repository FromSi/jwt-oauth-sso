package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewBaseLogoutDeviceRequest_And_NewBaseLogoutDeviceRequestBody_And_GetBody(t *testing.T) {
	requestBody := NewBaseLogoutDeviceRequestBody()

	assert.NotNil(t, requestBody)

	request := NewBaseLogoutDeviceRequest(requestBody)

	assert.NotNil(t, request)
	assert.Equal(t, requestBody, request.GetBody())
}

func TestNewBaseLogoutDeviceRequestBody_GetDeviceUUID(t *testing.T) {
	requestBody := NewBaseLogoutDeviceRequestBody()

	assert.NotNil(t, requestBody)

	requestBody.DeviceUUID = "1"

	assert.Equal(t, "1", requestBody.GetDeviceUUID())

	requestBody.DeviceUUID = "2"

	assert.Equal(t, "2", requestBody.GetDeviceUUID())
}

func TestNewBaseLogoutDeviceRequest_Make(t *testing.T) {
	requestConstructorBody := NewBaseLogoutDeviceRequestBody()

	assert.NotNil(t, requestConstructorBody)

	requestConstructor := NewBaseLogoutDeviceRequest(requestConstructorBody)

	assert.NotNil(t, requestConstructor)

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

			request, errResponse := requestConstructor.Make(c)

			if tt.error {
				assert.NotEmpty(t, errResponse)
				assert.Empty(t, request)
			} else {
				assert.Empty(t, errResponse)
				assert.NotEmpty(t, request)
				assert.NotEmpty(t, request.GetBody())
				assert.NotEmpty(t, request.GetBody().GetDeviceUUID())
			}
		})
	}
}
