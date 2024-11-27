package requests

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"strings"
)

type BearerAuthRequestHeader struct {
	AccessToken tokens.AccessToken
}

func NewBearerAuthRequestHeader(
	context *gin.Context,
	accessTokenBuilder tokens.AccessTokenBuilder,
) (*BearerAuthRequestHeader, error) {
	authHeader := context.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("authorization header missing")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	accessTokenBuilder, err := accessTokenBuilder.NewFromJwtString(token)

	if err != nil {
		return nil, err
	}

	accessToken, err := accessTokenBuilder.Build()

	if err != nil {
		return nil, err
	}

	return &BearerAuthRequestHeader{
		AccessToken: accessToken,
	}, nil
}
