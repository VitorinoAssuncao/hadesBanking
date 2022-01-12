package customError

import "errors"

var (
	ErrorTransferAccountNotFound = errors.New("account not found, please validate the data")
	ErrorTransferListing         = errors.New("error when listing all transfers")
)
