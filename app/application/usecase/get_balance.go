package usecase

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetBalance(ctx context.Context, accountID string) (output.AccountBalanceVO, error) {
	tempAccount, err := usecase.accountRepository.GetByID(ctx, types.AccountID(accountID))
	if err != nil {
		return output.AccountBalanceVO{}, errorAccountIDNotFound
	}
	return output.AccountBalanceVO{Balance: tempAccount.Balance.ToFloat()}, nil
}
