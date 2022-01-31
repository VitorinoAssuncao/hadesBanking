package account

import (
	"context"
	"errors"

	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetBalance(ctx context.Context, accountID string) (types.Money, error) {
	const operation = "Usecase.Account.GetBalance"

	balance, err := usecase.accountRepository.GetBalanceByAccountID(ctx, types.ExternalID(accountID))
	if err != nil {
		if errors.Is(err, customError.ErrorAccountIDNotFound) {
			usecase.logger.LogError(operation, err.Error())
			return -1, err
		}

		usecase.logger.LogError(operation, err.Error())
		return -1, customError.ErrorAccountIDSearching
	}

	usecase.logger.LogInfo(operation, "balance successfully listed")
	return balance, nil
}
