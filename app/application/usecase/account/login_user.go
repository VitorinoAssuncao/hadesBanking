package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) LoginUser(ctx context.Context, loginInput account.Account) (string, error) {
	const operation = "Usecase.Account.LoginUser"

	usecase.logRepository.LogInfo(operation, "trying to find the account")
	tempAccount, err := usecase.accountRepository.GetByCPF(context.Background(), loginInput.CPF.ToString())
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return "", customError.ErrorAccountLogin
	}

	usecase.logRepository.LogInfo(operation, "comparing the secret informed, and the secret in the database")
	err = tempAccount.Secret.CompareSecret(string(loginInput.Secret.ToString()))
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return "", customError.ErrorAccountLogin
	}

	usecase.logRepository.LogInfo(operation, "generating the authorization token")
	token, err := usecase.tokenRepository.GenerateToken(string(tempAccount.ExternalID))
	if err != nil {
		usecase.logRepository.LogError(operation, err.Error())
		return "", customError.ErrorAccountTokenGeneration
	}

	usecase.logRepository.LogInfo(operation, "account logged sucessfully")
	return token, nil
}
