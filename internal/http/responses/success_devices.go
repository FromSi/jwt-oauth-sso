package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

type SuccessDevicesResponse struct {
	Data []SuccessDevicesResponseData `json:"data"`
}

type SuccessDevicesResponseData struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"userUUID"`
	UserAgent string `json:"userAgent"`
	Ip        string `json:"ip"`
	IssuedAt  int    `json:"issuedAt"`
	ExpiresAt int    `json:"expiresAt"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

func NewSuccessDevicesResponse(devices []repositories.Device) *SuccessDevicesResponse {
	data := make([]SuccessDevicesResponseData, len(devices))

	for i, device := range devices {
		data[i] = SuccessDevicesResponseData{
			UUID:      device.GetUUID(),
			UserUUID:  device.GetUserUUID(),
			UserAgent: device.GetUserAgent(),
			Ip:        device.GetIp(),
			IssuedAt:  device.GetIssuedAt(),
			ExpiresAt: device.GetExpiresAt(),
			CreatedAt: device.GetCreatedAt(),
			UpdatedAt: device.GetUpdatedAt(),
		}
	}

	return &SuccessDevicesResponse{
		Data: data,
	}
}
