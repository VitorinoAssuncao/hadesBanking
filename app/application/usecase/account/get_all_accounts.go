package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) GetAll(ctx context.Context) ([]account.Account, error) {
	var accounts = make([]account.Account, 0)
	accounts, err := usecase.accountRepository.GetAll(ctx)
	if err != nil {
		return nil, customError.ErrorAccountsListing
	}

	return accounts, nil
}
