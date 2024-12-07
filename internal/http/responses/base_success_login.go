package responses

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
)

type BaseSuccessLoginResponse struct {
	Data struct {
		AuthType         string `json:"authType"`
		AccessToken      string `json:"accessToken"`
		RefreshToken     string `json:"refreshToken"`
		AccessExpiresIn  int    `json:"accessExpiresIn"`
		RefreshExpiresIn int    `json:"refreshExpiresIn"`
	} `json:"data"`
}

func NewBaseSuccessLoginResponse() *BaseSuccessLoginResponse {
	return &BaseSuccessLoginResponse{}
}

func (receiver BaseSuccessLoginResponse) Make(
	device repositories.Device,
) (SuccessLoginResponse, error) {
	accessToken, err := device.GenerateAccessToken()

	if err != nil {
		return nil, err
	}

	accessTokenToJWT, err := accessToken.ToString()

	if err != nil {
		return nil, err
	}

	return &BaseSuccessLoginResponse{
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
