package responses

//go:generate mockgen -destination=../../../mocks/http/responses/mock_error_internal_server_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses ErrorInternalServerResponse
type ErrorInternalServerResponse interface {
	Make(error) ErrorInternalServerResponse
}
