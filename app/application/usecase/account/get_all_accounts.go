package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) GetAll(ctx context.Context) ([]account.Account, error) {
	const operation = "Usecase.Account.GetAll"

	var accounts []account.Account
	accounts, err := usecase.accountRepository.GetAll(ctx)
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return nil, customError.ErrorAccountsListing
	}

	usecase.logRepository.LogInfo(operation, "listing data sucessfully")
	return accounts, nil
}
