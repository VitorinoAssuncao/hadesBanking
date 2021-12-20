package account

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/types"
)

func (usecase *usecase) GetBalance(ctx context.Context, accountID string) (output.AccountBalanceVO, error) {
	value, err := usecase.accountRepository.GetBalance(ctx, types.AccountID(accountID))
	if err != nil {
		return output.AccountBalanceVO{Balance: 0}, err
	}
	return output.AccountBalanceVO{Balance: value.ToFloat()}, nil
}
