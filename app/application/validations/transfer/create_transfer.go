package validations

import (
	"stoneBanking/app/application/vo/input"
)

func ValidateTransferData(transferData input.CreateTransferVO) (input.CreateTransferVO, error) {
	if transferData.AccountOrigin_ID == transferData.AccountDestiny_ID {
		return input.CreateTransferVO{}, ErrorTransferOriginEqualDestiny
	}

	if transferData.AccountOrigin_ID == "" {
		return input.CreateTransferVO{}, ErrorTransferAccountOriginIDRequired
	}

	if transferData.AccountDestiny_ID == "" {
		return input.CreateTransferVO{}, ErrorTransferAccountDestinyIDRequired
	}

	if transferData.Amount <= 0 {
		return input.CreateTransferVO{}, ErrorTransferAmountInvalid
	}

	return input.CreateTransferVO{}, nil
}
