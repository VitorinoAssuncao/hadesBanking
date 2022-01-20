package account

import (
	"context"
	"errors"
	validations "stoneBanking/app/application/usecase/account/validations"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) Create(ctx context.Context, accountData account.Account) (account.Account, error) {
	const operation = "Usecase.Account.Create"

	usecase.logRepository.LogInfo(operation, "begin validation of the received data")
	err := validations.ValidateAccountData(accountData)
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return account.Account{}, err
	}

	_, err = usecase.accountRepository.GetByCPF(ctx, accountData.CPF.ToString())
	//validate if account with that cpf exist, if not continue the creation of a new account
	if err == nil {
		return account.Account{}, customError.ErrorAccountCPFExists
	}

	if !errors.Is(err, customError.ErrorAccountCPFNotFound) {
		usecase.logRepository.LogError(operation, err.Error())
		return account.Account{}, customError.ErrorCreateAccount
	}

	usecase.logRepository.LogInfo(operation, "create the account in database")
	accountResult, err := usecase.accountRepository.Create(ctx, accountData)

	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return account.Account{}, customError.ErrorCreateAccount
	}

	usecase.logRepository.LogInfo(operation, "account created sucessfully")
	return accountResult, nil
}
