package responses

type ErrorConflictResponse struct {
	Message string `json:"message"`
}

func NewErrorConflictResponse(err error) *ErrorConflictResponse {
	return &ErrorConflictResponse{
		Message: err.Error(),
	}
}
