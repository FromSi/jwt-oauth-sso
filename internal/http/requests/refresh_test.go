package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewRefreshRequestBody(t *testing.T) {
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

			requestBody, err := NewRefreshRequestBody(c)

			if tt.error {
				assert.Error(t, err)
				assert.Nil(t, requestBody)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, requestBody)
			}
		})
	}
}
