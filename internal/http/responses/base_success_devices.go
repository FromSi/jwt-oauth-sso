package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

type BaseSuccessDevicesResponse struct {
	Data []BaseSuccessDevicesResponseData `json:"data"`
}

type BaseSuccessDevicesResponseData struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"userUUID"`
	UserAgent string `json:"userAgent"`
	Ip        string `json:"ip"`
	IssuedAt  int    `json:"issuedAt"`
	ExpiresAt int    `json:"expiresAt"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

func NewBaseSuccessDevicesResponse() *BaseSuccessDevicesResponse {
	return &BaseSuccessDevicesResponse{}
}

func (receiver BaseSuccessDevicesResponse) Make(
	devices []repositories.Device,
) SuccessDevicesResponse {
	data := make([]BaseSuccessDevicesResponseData, len(devices))

	for i, device := range devices {
		data[i] = BaseSuccessDevicesResponseData{
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

	return &BaseSuccessDevicesResponse{
		Data: data,
	}
}
