package account

import (
	"context"
	"stoneBanking/app/common/utils"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (usecase *usecase) LoginUser(ctx context.Context, loginInput account.Account) (string, error) {

	tempAccount, err := usecase.accountRepository.GetByCPF(context.Background(), loginInput.CPF)
	if err != nil {
		return "", customError.ErrorAccountLogin
	}

	if !utils.ValidateHash(tempAccount.Secret, loginInput.Secret) {
		return "", customError.ErrorAccountLogin
	}

	token, err := utils.GenerateToken(string(tempAccount.ExternalID))
	if err != nil {
		return "", customError.ErrorAccountTokenGeneration
	}

	return token, nil
}
