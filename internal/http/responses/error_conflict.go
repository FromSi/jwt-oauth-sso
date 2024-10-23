package responses

type ErrorConflictResponse struct {
	Message string `json:"message"`
}

func NewErrorConflictResponse(message string) *ErrorConflictResponse {
	return &ErrorConflictResponse{
		Message: message,
	}
}
