package transfer

import (
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/transfer/vo/input"
)

func ValidateTransferData(transferData input.CreateTransferVO) (input.CreateTransferVO, error) {
	if transferData.AccountOriginID == transferData.AccountDestinyID {
		return input.CreateTransferVO{}, customError.ErrorTransferOriginEqualDestiny
	}

	if transferData.AccountOriginID == "" {
		return input.CreateTransferVO{}, customError.ErrorTransferAccountOriginIDRequired
	}

	if transferData.AccountDestinyID == "" {
		return input.CreateTransferVO{}, customError.ErrorTransferAccountDestinyIDRequired
	}

	if transferData.Amount <= 0 {
		return input.CreateTransferVO{}, customError.ErrorTransferAmountInvalid
	}

	return transferData, nil
}
