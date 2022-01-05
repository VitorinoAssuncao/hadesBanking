package transfer

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetAllByID(ctx context.Context, accountID types.AccountExternalID) ([]output.TransferOutputVO, error) {
	var resultTransfers = make([]output.TransferOutputVO, 0)

	account, err := usecase.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		return resultTransfers, err
	}

	transfers, err := usecase.transferRepository.GetAllByAccountID(ctx, types.InternalID(account.ID))
	if err != nil {
		return resultTransfers, err
	}

	for _, transfer := range transfers {
		transferOutput := output.TransferToTransferOutput(transfer)
		resultTransfers = append(resultTransfers, transferOutput)
	}

	return resultTransfers, nil
}
