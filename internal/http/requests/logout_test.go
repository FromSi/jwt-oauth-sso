package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewLogoutRequest_And_NewLogoutRequestBody(t *testing.T) {
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

			request, errResponse := NewLogoutRequest(c)

			if tt.error {
				assert.NotNil(t, errResponse)
				assert.Nil(t, request)
			} else {
				assert.Nil(t, errResponse)
				assert.NotNil(t, request)
				assert.NotNil(t, request.Body)
			}
		})
	}
}
