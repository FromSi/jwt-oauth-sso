package responses

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
)

type SuccessLoginResponse struct {
	Data struct {
		AuthType         string `json:"authType"`
		AccessToken      string `json:"accessToken"`
		RefreshToken     string `json:"refreshToken"`
		AccessExpiresIn  int    `json:"accessExpiresIn"`
		RefreshExpiresIn int    `json:"refreshExpiresIn"`
	} `json:"data"`
}

func NewSuccessLoginResponse(
	configs configs.TokenConfig,
	device repositories.Device,
) (*SuccessLoginResponse, error) {
	accessToken, err := device.GenerateAccessToken(configs)

	if err != nil {
		return nil, err
	}

	accessTokenToJWT, err := accessToken.GetJWT()

	if err != nil {
		return nil, err
	}

	return &SuccessLoginResponse{
		Data: struct {
			AuthType         string `json:"authType"`
			AccessToken      string `json:"accessToken"`
			RefreshToken     string `json:"refreshToken"`
			AccessExpiresIn  int    `json:"accessExpiresIn"`
			RefreshExpiresIn int    `json:"refreshExpiresIn"`
		}{
			AuthType:         "bearer",
			AccessToken:      accessTokenToJWT,
			RefreshToken:     device.GetRefreshToken(),
			AccessExpiresIn:  int(accessToken.ExpirationTime.Unix()),
			RefreshExpiresIn: device.GetExpiresAt(),
		},
	}, err
}
