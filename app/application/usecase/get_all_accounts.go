package usecase

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
)

func (usecase *usecase) GetAll(ctx context.Context) ([]output.AccountOutputVO, error) {
	var accounts = make([]account.Account, 0)
	var resultAccounts = make([]output.AccountOutputVO, 0)
	accounts, err := usecase.accountRepository.GetAll(ctx)
	for _, account := range accounts {
		accountOutput := output.AccountToOutput(account)
		resultAccounts = append(resultAccounts, accountOutput)
	}

	if err != nil {
		return resultAccounts, err
	}

	return resultAccounts, nil
}
