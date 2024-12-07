package responses

type BaseErrorInternalServerResponse struct {
	Message string `json:"message"`
}

func NewBaseErrorInternalServerResponse() *BaseErrorInternalServerResponse {
	return &BaseErrorInternalServerResponse{}
}

func (receiver BaseErrorInternalServerResponse) Make(
	err error,
) ErrorInternalServerResponse {
	return &BaseErrorInternalServerResponse{
		Message: err.Error(),
	}
}
