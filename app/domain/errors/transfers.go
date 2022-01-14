package customError

import "errors"

var (
	ErrorTransferAccountNotFound         = errors.New("account not found, please validate the data")
	ErrorTransferListing                 = errors.New("error when listing all transfers")
	ErrorTransferCreateOriginError       = errors.New("account origin not found, please validate")
	ErrorTransferCreateDestinyError      = errors.New("account destination not found, please validate")
	ErrorTransferCreateInsufficientFunds = errors.New("insuficient funds")
	ErrorTransferCreate                  = errors.New("error when creating transfer, please try again")
)