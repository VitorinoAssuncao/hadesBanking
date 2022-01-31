package response

type OutputError struct {
	Error string `json:"error" example:"error"`
}

func NewError(err error) *OutputError {
	return &OutputError{
		Error: err.Error(),
	}
}
