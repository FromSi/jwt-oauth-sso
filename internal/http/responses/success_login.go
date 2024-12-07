package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

//go:generate mockgen -destination=../../../mocks/http/responses/mock_success_login_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses SuccessLoginResponse
type SuccessLoginResponse interface {
	Make(repositories.Device) (SuccessLoginResponse, error)
}
