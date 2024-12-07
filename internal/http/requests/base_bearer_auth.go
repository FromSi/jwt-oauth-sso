package requests

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/tokens"
	"github.com/gin-gonic/gin"
	"strings"
)

type BaseBearerAuthRequestHeader struct {
	AccessToken        tokens.AccessToken
	AccessTokenBuilder tokens.AccessTokenBuilder
}

func NewBaseBearerAuthRequestHeader(
	accessTokenBuilder tokens.AccessTokenBuilder,
) *BaseBearerAuthRequestHeader {
	return &BaseBearerAuthRequestHeader{
		AccessTokenBuilder: accessTokenBuilder,
	}
}

func (receiver BaseBearerAuthRequestHeader) Make(
	context *gin.Context,
) (BearerAuthRequestHeader, error) {
	authHeader := context.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("authorization header missing")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	accessTokenBuilder, err := receiver.AccessTokenBuilder.NewFromJwtString(token)

	if err != nil {
		return nil, err
	}

	accessToken, err := accessTokenBuilder.Build()

	if err != nil {
		return nil, err
	}

	return &BaseBearerAuthRequestHeader{
		AccessToken:        accessToken,
		AccessTokenBuilder: receiver.AccessTokenBuilder,
	}, nil
}

func (receiver BaseBearerAuthRequestHeader) GetAccessToken() tokens.AccessToken {
	return receiver.AccessToken
}
