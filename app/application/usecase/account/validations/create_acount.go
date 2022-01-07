package account

import (
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func ValidateAccountData(accountData account.Account) error {
	if accountData.Name == "" {
		return customError.ErrorAccountNameRequired
	}

	if accountData.CPF == "" {
		return customError.ErrorAccountCPFRequired
	}

	if accountData.Secret == "" {
		return customError.ErrorAccountSecretRequired
	}

	if accountData.Balance < 0 {
		return customError.ErrorAccountBalanceInvalid
	}

	return nil
}
