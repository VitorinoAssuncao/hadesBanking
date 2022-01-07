package transfer

import "errors"

var (
	ErrorTransferOriginEqualDestiny       = errors.New("conta de origem e destino são iguais, favor validar os id's informados")
	ErrorTransferAccountOriginIDRequired  = errors.New("campo conta de origem é obrigatório, e deve ser informado diferente de vazio")
	ErrorTransferAccountDestinyIDRequired = errors.New("campo conta de destino é obrigatório, e deve ser informado diferente de vazio")
	ErrorTransferAmountInvalid            = errors.New("valor da transferencia não pode ser igual ou menor que zero (0)")
)
