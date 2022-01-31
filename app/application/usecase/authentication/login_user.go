package authentication

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) LoginUser(ctx context.Context, loginInput account.Account) (string, error) {
	const operation = "Usecase.Account.LoginUser"

	usecase.logger.LogInfo(operation, "trying to find the account")
	tempAccount, err := usecase.accountRepository.GetCredentialByCPF(context.Background(), loginInput.CPF.ToString())
	if err != nil {
		usecase.logger.LogError(operation, err.Error())
		return "", customError.ErrorAccountLogin
	}

	usecase.logger.LogInfo(operation, "comparing the secret informed, and the secret in the database")
	err = tempAccount.Secret.CompareSecret(string(loginInput.Secret.ToString()))
	if err != nil {
		usecase.logger.LogError(operation, err.Error())
		return "", customError.ErrorAccountLogin
	}

	usecase.logger.LogInfo(operation, "generating the authorization token")
	token, err := usecase.token.GenerateToken(string(tempAccount.ExternalID))
	if err != nil {
		usecase.logger.LogError(operation, err.Error())
		return "", customError.ErrorAccountTokenGeneration
	}

	usecase.logger.LogInfo(operation, "account logged successfully")
	return token, nil
}
