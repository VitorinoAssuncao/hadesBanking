package account

import (
	"context"
	validations "stoneBanking/app/application/usecase/account/validations"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) Create(ctx context.Context, accountData account.Account) (account.Account, error) {
	err := validations.ValidateAccountData(accountData)
	if err != nil {
		return account.Account{}, err
	}
	_, err = usecase.accountRepository.GetByCPF(ctx, accountData.CPF)
	//validate if account with that cpf exist, if not continue the creation of a new account
	if err == nil {
		return account.Account{}, customError.ErrorAccountCPFExists
	}

	accountResult, err := usecase.accountRepository.Create(ctx, accountData)

	if err != nil {
		return account.Account{}, customError.ErrorCreateAccount
	}

	return accountResult, nil
}
