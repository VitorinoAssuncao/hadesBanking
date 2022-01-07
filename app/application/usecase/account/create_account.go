package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (usecase *usecase) Create(ctx context.Context, accountData account.Account) (account.Account, error) {

	_, err := usecase.accountRepository.GetByCPF(ctx, accountData.CPF)
	//validate if account with that cpf exist, if not continue the creation of a new account
	if err == nil {
		return account.Account{}, ErrorAccountCPFExists
	}

	accountResult, err := usecase.accountRepository.Create(ctx, accountData)

	if err != nil {
		return account.Account{}, ErrorCreateAccount
	}

	return accountResult, nil
}
