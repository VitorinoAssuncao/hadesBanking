package account

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
)

func (usecase *usecase) GetAll(ctx context.Context) ([]output.AccountOutputVO, error) {
	var accounts = make([]account.Account, 0)
	var resultAccounts = make([]output.AccountOutputVO, 0)
	accounts, err := usecase.accountRepository.GetAll(ctx)
	if err != nil {
		return resultAccounts, ErrorAccountsListing
	}

	for _, account := range accounts {
		accountOutput := output.AccountToOutput(account)
		resultAccounts = append(resultAccounts, accountOutput)
	}

	return resultAccounts, nil
}