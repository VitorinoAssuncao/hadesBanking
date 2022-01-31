package transfer

import (
	"context"
	"errors"

	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]transfer.Transfer, error) {
	const operation = "Usecase.Transfer.GetAllByAccountID"

	usecase.logger.LogInfo(operation, "searching if the account exist")
	account, err := usecase.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		usecase.logger.LogError(operation, err.Error())
		return []transfer.Transfer{}, customError.ErrorTransferAccountNotFound
	}

	usecase.logger.LogInfo(operation, "searching all the transfers that the accounts is involved")
	transfers, err := usecase.transferRepository.GetAllByAccountID(ctx, types.InternalID(account.ID))
	if err != nil {
		if errors.Is(err, customError.ErrorTransferAccountNotFound) {
			usecase.logger.LogError(operation, err.Error())
			return []transfer.Transfer{}, err
		}

		usecase.logger.LogError(operation, err.Error())
		return []transfer.Transfer{}, customError.ErrorTransferListing
	}

	usecase.logger.LogInfo(operation, "listing all the data successfully")
	return transfers, nil
}
