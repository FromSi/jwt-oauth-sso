package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

//go:generate mockgen -destination=../../../mocks/http/responses/mock_success_register_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses SuccessRegisterResponse
type SuccessRegisterResponse interface {
	Make(repositories.Device) (SuccessRegisterResponse, error)
}
