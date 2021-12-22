package usecase

import (
	"context"
	"stoneBanking/app/application/validations"
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

	if err == nil {
		accountOutput = output.AccountToOutput(tempAccount)
		return &accountOutput, errorAccountCPFExists
	}

	account := input.GenerateAccount(accountData)

	accountResult, err := usecase.accountRepository.Create(ctx, account)

	if err != nil {
		return &accountOutput, errorCreateAccount
	}

	accountOutput = output.AccountToOutput(accountResult)

	return &accountOutput, err
}
