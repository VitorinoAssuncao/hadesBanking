package customError

import "errors"

var (
	ErrorTransferAccountNotFound          = errors.New("account not found, please validate the data")
	ErrorTransferListing                  = errors.New("error when listing all transfers")
	ErrorTransferCreateOriginError        = errors.New("account origin not found, please validate")
	ErrorTransferCreateDestinyError       = errors.New("account destination not found, please validate")
	ErrorTransferCreateInsufficientFunds  = errors.New("insuficient funds")
	ErrorTransferCreate                   = errors.New("error when creating transfer, please try again")
	ErrorTransferOriginEqualDestiny       = errors.New("origin account and destiny account cannot by the same")
	ErrorTransferAccountOriginIDRequired  = errors.New("field account_origin_id is required")
	ErrorTransferAccountDestinyIDRequired = errors.New("field account_destination_id is required")
	ErrorTransferAmountInvalid            = errors.New("transfer amount should be greater than 0")
)
