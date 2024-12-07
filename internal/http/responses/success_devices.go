package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

//go:generate mockgen -destination=../../../mocks/http/responses/mock_success_devices_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses SuccessDevicesResponse
type SuccessDevicesResponse interface {
	Make([]repositories.Device) SuccessDevicesResponse
}
