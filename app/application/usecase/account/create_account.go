package account

import (
	"context"
	"errors"

	validations "stoneBanking/app/application/usecase/account/validations"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (u *usecase) Create(ctx context.Context, accountData account.Account) (account.Account, error) {
	const operation = "Usecase.Account.Create"

	u.logger.LogDebug(operation, "begin validation of the received data")
	err := validations.ValidateAccountData(accountData)
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return account.Account{}, err
	}

	_, err = u.accountRepository.GetCredentialByCPF(ctx, accountData.CPF.ToString())
	//validate if account with that cpf exist, if not continue the creation of a new account
	if err == nil {
		u.logger.LogError(operation, customError.ErrorAccountCPFExists.Error())
		return account.Account{}, customError.ErrorAccountCPFExists
	}

	if !errors.Is(err, customError.ErrorAccountCPFNotFound) {
		u.logger.LogError(operation, err.Error())
		return account.Account{}, customError.ErrorCreateAccount
	}

	u.logger.LogDebug(operation, "create the account in database")
	accountResult, err := u.accountRepository.Create(ctx, accountData)

	if err != nil {
		u.logger.LogError(operation, err.Error())
		return account.Account{}, customError.ErrorCreateAccount
	}

	u.logger.LogDebug(operation, "account created successfully")
	return accountResult, nil
}
