package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

type SuccessDevicesResponse struct {
	Data []SuccessDevicesResponseData `json:"data"`
}

type SuccessDevicesResponseData struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"userUUID"`
	Agent     string `json:"agent"`
	Ip        string `json:"ip"`
	ExpiredAt int    `json:"expiredAt"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

func NewSuccessDevicesResponse(devices []repositories.Device) *SuccessDevicesResponse {
	data := make([]SuccessDevicesResponseData, len(devices))

	for i, device := range devices {
		data[i] = SuccessDevicesResponseData{
			UUID:      device.GetUUID(),
			UserUUID:  device.GetUserUUID(),
			Agent:     device.GetAgent(),
			Ip:        device.GetIp(),
			ExpiredAt: device.GetExpiredAt(),
			CreatedAt: device.GetCreatedAt(),
			UpdatedAt: device.GetUpdatedAt(),
		}
	}

	return &SuccessDevicesResponse{
		Data: data,
	}
}
