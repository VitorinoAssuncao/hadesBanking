package usecase

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetBalance(ctx context.Context, accountID string) (output.AccountBalanceVO, error) {
	var tempAccount = &account.Account{}
	tempAccount, err := usecase.accountRepository.GetByID(ctx, types.AccountID(accountID), tempAccount)
	if err != nil {
		return output.AccountBalanceVO{Balance: 0}, err
	}
	return output.AccountBalanceVO{Balance: tempAccount.Balance.ToFloat()}, nil
}
