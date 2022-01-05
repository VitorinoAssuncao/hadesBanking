package account

import (
	"context"
	validations "stoneBanking/app/application/validations/account"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
)

func (usecase *usecase) Create(ctx context.Context, accountData input.CreateAccountVO) (*output.AccountOutputVO, error) {
	var accountOutput output.AccountOutputVO
	var err error
	accountData, err = validations.ValidateAccountInput(accountData)

	if err != nil {
		return &accountOutput, err
	}

	tempAccount, err := usecase.accountRepository.GetByCPF(ctx, accountData.CPF)
	//validate if account with that cpf exist, if not continue the creation of a new account
	if err == nil {
		accountOutput = output.AccountToOutput(tempAccount)
		return &output.AccountOutputVO{}, ErrorAccountCPFExists
	}

	account := input.GenerateAccount(accountData)

	accountResult, err := usecase.accountRepository.Create(ctx, account)

	if err != nil {
		return &output.AccountOutputVO{}, ErrorCreateAccount
	}

	accountOutput = output.AccountToOutput(accountResult)

	return &accountOutput, nil
}
