package responses

type ErrorBadRequestResponse struct {
	Errors map[string][]string `json:"errors"`
}

func NewErrorBadRequestResponse(errors map[string][]string) *ErrorBadRequestResponse {
	return &ErrorBadRequestResponse{
		Errors: errors,
	}
}
