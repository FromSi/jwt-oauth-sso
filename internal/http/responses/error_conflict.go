package responses

//go:generate mockgen -destination=../../../mocks/http/responses/mock_error_conflict_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses ErrorConflictResponse
type ErrorConflictResponse interface {
	Make(error) ErrorConflictResponse
}
