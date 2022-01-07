package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]transfer.Transfer, error) {
	var resultTransfers = make([]transfer.Transfer, 0)

	account, err := usecase.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		return resultTransfers, err
	}

	transfers, err := usecase.transferRepository.GetAllByAccountID(ctx, types.InternalID(account.ID))
	if err != nil {
		return resultTransfers, err
	}

	return transfers, nil
}
