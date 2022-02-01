package account

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (u *usecase) GetAll(ctx context.Context) ([]account.Account, error) {
	const operation = "Usecase.Account.GetAll"

	var accounts []account.Account
	accounts, err := u.accountRepository.GetAll(ctx)
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return nil, customError.ErrorAccountsListing
	}

	u.logger.LogInfo(operation, "listing data successfully")
	return accounts, nil
}
