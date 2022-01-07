package account

import (
	"context"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetBalance(ctx context.Context, accountID string) (float64, error) {
	tempAccount, err := usecase.accountRepository.GetByID(ctx, types.ExternalID(accountID))
	if err != nil {
		return -1, customError.ErrorAccountIDNotFound
	}
	return tempAccount.Balance.ToFloat(), nil
}
