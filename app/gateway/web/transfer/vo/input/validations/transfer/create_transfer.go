package transfer

import (
	"stoneBanking/app/gateway/web/transfer/vo/input"
)

func ValidateTransferData(transferData input.CreateTransferVO) (input.CreateTransferVO, error) {
	if transferData.AccountOriginID == transferData.AccountDestinyID {
		return input.CreateTransferVO{}, ErrorTransferOriginEqualDestiny
	}

	if transferData.AccountOriginID == "" {
		return input.CreateTransferVO{}, ErrorTransferAccountOriginIDRequired
	}

	if transferData.AccountDestinyID == "" {
		return input.CreateTransferVO{}, ErrorTransferAccountDestinyIDRequired
	}

	if transferData.Amount <= 0 {
		return input.CreateTransferVO{}, ErrorTransferAmountInvalid
	}

	return transferData, nil
}
