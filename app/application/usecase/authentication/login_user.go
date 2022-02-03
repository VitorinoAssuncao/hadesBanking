package authentication

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (u *usecase) LoginUser(ctx context.Context, loginInput account.Account) (string, error) {
	const operation = "Usecase.Account.LoginUser"

	u.logger.LogDebug(operation, "trying to find the account")
	tempAccount, err := u.accountRepository.GetCredentialByCPF(context.Background(), loginInput.CPF.ToString())
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return "", customError.ErrorAccountLogin
	}

	u.logger.LogDebug(operation, "comparing the secret informed, and the secret in the database")
	err = tempAccount.Secret.CompareSecret(string(loginInput.Secret.ToString()))
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return "", customError.ErrorAccountLogin
	}

	u.logger.LogDebug(operation, "generating the authorization token")
	token, err := u.token.GenerateToken(string(tempAccount.ExternalID))
	if err != nil {
		u.logger.LogError(operation, err.Error())
		return "", customError.ErrorAccountTokenGeneration
	}

	u.logger.LogDebug(operation, "account logged successfully")
	return token, nil
}
