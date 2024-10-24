package responses

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
)

type SuccessRefreshResponse struct {
	Data struct {
		AuthType         string `json:"authType"`
		AccessToken      string `json:"accessToken"`
		RefreshToken     string `json:"refreshToken"`
		AccessExpiresIn  int    `json:"accessExpiresIn"`
		RefreshExpiresIn int    `json:"refreshExpiresIn"`
	} `json:"data"`
}

func NewSuccessRefreshResponse(configs configs.TokenConfig, device repositories.Device) (*SuccessRefreshResponse, error) {
	accessToken, err := device.GenerateAccessToken(configs)

	if err != nil {
		return nil, err
	}

	accessTokenToJWT, err := accessToken.GetJWT()

	if err != nil {
		return nil, err
	}

	return &SuccessRefreshResponse{
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
			RefreshExpiresIn: device.GetExpiredAt(),
		},
	}, err
}
