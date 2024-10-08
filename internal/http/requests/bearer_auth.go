package requests

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"strings"
)

type BearerAuthRequestHeader struct {
	AccessToken tokens.AccessToken
}

func NewBearerAuthRequestHeader(context *gin.Context, config configs.TokenConfig) (*BearerAuthRequestHeader, error) {
	authHeader := context.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("authorization header missing")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	accessToken, err := tokens.NewAccessTokenByJWT(config, token)

	if err != nil {
		return nil, err
	}

	return &BearerAuthRequestHeader{
		AccessToken: *accessToken,
	}, nil
}
