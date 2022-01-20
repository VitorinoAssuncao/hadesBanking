package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]transfer.Transfer, error) {
	const operation = "Usecase.Transfer.GetAllByAccountID"

	usecase.logRepository.LogInfo(operation, "searching if the account exist")
	account, err := usecase.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return []transfer.Transfer{}, customError.ErrorTransferAccountNotFound
	}

	usecase.logRepository.LogInfo(operation, "searching all the transfers that the accounts is involved")
	transfers, err := usecase.transferRepository.GetAllByAccountID(ctx, types.InternalID(account.ID))
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return []transfer.Transfer{}, customError.ErrorTransferListing
	}

	usecase.logRepository.LogInfo(operation, "listing all the data sucessfully")
	return transfers, nil
}
