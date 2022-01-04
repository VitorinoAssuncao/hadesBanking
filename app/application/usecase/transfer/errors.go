package transfer

import "errors"

var (
	ErrorTransferCreateOriginError       = errors.New("não foi possível realizar a transferencia, favor validar o id informado")
	ErrorTransferCreateDestinyError      = errors.New("não foi possível realizar a transferencia, favor validar o id informado")
	ErrorTransferCreateInsufficientFunds = errors.New("o saldo na conta de origem não é o bastante para essa transação")
)
