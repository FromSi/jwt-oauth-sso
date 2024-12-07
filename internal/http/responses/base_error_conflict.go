package responses

type BaseErrorConflictResponse struct {
	Message string `json:"message"`
}

func NewBaseErrorConflictResponse() *BaseErrorConflictResponse {
	return &BaseErrorConflictResponse{}
}

func (receiver BaseErrorConflictResponse) Make(
	err error,
) ErrorConflictResponse {
	return &BaseErrorConflictResponse{
		Message: err.Error(),
	}
}
