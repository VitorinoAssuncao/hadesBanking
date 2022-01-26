package account

import (
	"context"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetBalance(ctx context.Context, accountID string) (float64, error) {
	const operation = "Usecase.Account.GetBalance"

	tempAccount, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(accountID))
	if err != nil {
		usecase.logger.LogError(operation, err.Error())
		return -1, customError.ErrorAccountIDNotFound
	}

	usecase.logger.LogInfo(operation, "balance sucessfully listed")
	return tempAccount.Balance.ToFloat(), nil
}
