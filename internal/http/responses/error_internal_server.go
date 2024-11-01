package responses

type ErrorInternalServerResponse struct {
	Message string `json:"message"`
}

func NewErrorInternalServerResponse(err error) *ErrorInternalServerResponse {
	return &ErrorInternalServerResponse{
		Message: err.Error(),
	}
}
