package response

type OutputError struct {
	Error string `json:"error" example:"error"`
}

func NewError(err error) *OutputError {
	return &OutputError{
		Error: err.Error(),
	}
}

func NewErrors(errs []error) *[]OutputError {
	var output = make([]OutputError, 0)
	for _, err := range errs {
		output = append(output, OutputError{Error: err.Error()})
	}
	return &output
}
