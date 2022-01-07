package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (usecase *usecase) GetAll(ctx context.Context) ([]account.Account, error) {
	var accounts = make([]account.Account, 0)
	accounts, err := usecase.accountRepository.GetAll(ctx)
	if err != nil {
		return accounts, ErrorAccountsListing
	}

	return accounts, nil
}
