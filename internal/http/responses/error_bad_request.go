package responses

//go:generate mockgen -destination=../../../mocks/http/responses/mock_error_bad_request_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses ErrorBadRequestResponse
type ErrorBadRequestResponse interface {
	Make(error) ErrorBadRequestResponse
}
