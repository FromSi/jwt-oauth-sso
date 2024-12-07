package responses

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

//go:generate mockgen -destination=../../../mocks/http/responses/mock_success_refresh_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses SuccessRefreshResponse
type SuccessRefreshResponse interface {
	Make(repositories.Device) (SuccessRefreshResponse, error)
}
