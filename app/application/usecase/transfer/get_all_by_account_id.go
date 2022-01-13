package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]transfer.Transfer, error) {
	account, err := usecase.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		return []transfer.Transfer{}, customError.ErrorTransferAccountNotFound
	}

	transfers, err := usecase.transferRepository.GetAllByAccountID(ctx, types.InternalID(account.ID))
	if err != nil {
		return []transfer.Transfer{}, customError.ErrorTransferListing
	}

	return transfers, nil
}
