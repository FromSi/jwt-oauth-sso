package requests

import (
	"errors"
	tokens_mocks "github.com/fromsi/jwt-oauth-sso/mocks/tokens"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_NewBearerAuthRequestHeader(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockAccessToken := tokens_mocks.NewMockAccessToken(mockController)
	mockAccessTokenBuilder := tokens_mocks.NewMockAccessTokenBuilder(mockController)

	tests := []struct {
		name                    string
		accessToken             string
		errorEmptyHeader        bool
		errorAccessTokenBuilder error
		errorAccessToken        error
	}{
		{
			name:                    "Valid request",
			accessToken:             "1",
			errorEmptyHeader:        false,
			errorAccessTokenBuilder: nil,
			errorAccessToken:        nil,
		},
		{
			name:                    "Empty header",
			accessToken:             "",
			errorEmptyHeader:        true,
			errorAccessTokenBuilder: nil,
			errorAccessToken:        nil,
		},
		{
			name:                    "Not valid access token builder",
			accessToken:             "2",
			errorEmptyHeader:        false,
			errorAccessTokenBuilder: errors.New("error"),
			errorAccessToken:        nil,
		},
		{
			name:                    "Not valid access token",
			accessToken:             "2",
			errorEmptyHeader:        false,
			errorAccessTokenBuilder: nil,
			errorAccessToken:        errors.New("error"),
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest("POST", "", strings.NewReader(""))

			if len(tt.accessToken) != 0 {
				c.Request.Header.Set("Authorization", "Bearer "+tt.accessToken)
			}

			if !tt.errorEmptyHeader && tt.errorAccessTokenBuilder == nil {
				mockAccessTokenBuilder.EXPECT().NewFromJwtString(tt.accessToken).Return(mockAccessTokenBuilder, nil)
			} else if tt.errorAccessTokenBuilder != nil {
				mockAccessTokenBuilder.EXPECT().NewFromJwtString(tt.accessToken).Return(nil, tt.errorAccessTokenBuilder)
			}

			if !tt.errorEmptyHeader && tt.errorAccessTokenBuilder == nil && tt.errorAccessToken == nil {
				mockAccessTokenBuilder.EXPECT().Build().Return(mockAccessToken, nil)
			} else if tt.errorAccessToken != nil {
				mockAccessTokenBuilder.EXPECT().Build().Return(nil, tt.errorAccessToken)
			}

			request, err := NewBearerAuthRequestHeader(c, mockAccessTokenBuilder)

			if tt.errorEmptyHeader || tt.errorAccessTokenBuilder != nil || tt.errorAccessToken != nil {
				assert.Error(t, err)
				assert.Empty(t, request)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, request)
			}
		})
	}
}
