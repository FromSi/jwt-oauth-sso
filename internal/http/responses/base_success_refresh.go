package responses

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
)

type BaseSuccessRefreshResponse struct {
	Data struct {
		AuthType         string `json:"authType"`
		AccessToken      string `json:"accessToken"`
		RefreshToken     string `json:"refreshToken"`
		AccessExpiresIn  int    `json:"accessExpiresIn"`
		RefreshExpiresIn int    `json:"refreshExpiresIn"`
	} `json:"data"`
}

func NewBaseSuccessRefreshResponse() *BaseSuccessRefreshResponse {
	return &BaseSuccessRefreshResponse{}
}

func (receiver BaseSuccessRefreshResponse) Make(
	device repositories.Device,
) (SuccessRefreshResponse, error) {
	accessToken, err := device.GenerateAccessToken()

	if err != nil {
		return nil, err
	}

	accessTokenToJWT, err := accessToken.ToString()

	if err != nil {
		return nil, err
	}

	return &BaseSuccessRefreshResponse{
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
			AccessExpiresIn:  accessToken.GetExpirationTime(),
			RefreshExpiresIn: device.GetExpiresAt(),
		},
	}, err
}