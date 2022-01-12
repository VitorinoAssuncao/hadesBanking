package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) LoginUser(ctx context.Context, loginInput account.Account) (string, error) {
	tempAccount, err := usecase.accountRepository.GetByCPF(context.Background(), loginInput.CPF.ToString())
	if err != nil {
		return "", customError.ErrorAccountLogin
	}

	err = tempAccount.Secret.CompareSecret(string(loginInput.Secret.ToString()))
	if err != nil {
		return "", customError.ErrorAccountLogin
	}

	token, err := usecase.tokenRepository.GenerateToken(string(tempAccount.ExternalID))
	if err != nil {
		return "", customError.ErrorAccountTokenGeneration
	}

	return token, nil
}
