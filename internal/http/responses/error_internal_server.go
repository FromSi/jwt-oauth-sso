package responses

type ErrorInternalServerResponse struct {
	Message string `json:"message"`
}

func NewErrorInternalServerResponse(message string) *ErrorInternalServerResponse {
	return &ErrorInternalServerResponse{
		Message: message,
	}
}
