package transfer

import (
	"context"
	"errors"

	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (u *usecase) GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]transfer.Transfer, error) {
	const operation = "Usecase.Transfer.GetAllByAccountID"

	u.logger.LogInfo(operation, "searching if the account exist")
	account, err := u.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return []transfer.Transfer{}, customError.ErrorTransferAccountNotFound
	}

	u.logger.LogInfo(operation, "searching all the transfers that the accounts is involved")
	transfers, err := u.transferRepository.GetAllByAccountID(ctx, types.InternalID(account.ID))
	if err != nil {
		if errors.Is(err, customError.ErrorTransferAccountNotFound) {
			u.logger.LogError(operation, err.Error())
			return []transfer.Transfer{}, err
		}

		u.logger.LogError(operation, err.Error())
		return []transfer.Transfer{}, customError.ErrorTransferListing
	}

	u.logger.LogInfo(operation, "listing all the data successfully")
	return transfers, nil
}
