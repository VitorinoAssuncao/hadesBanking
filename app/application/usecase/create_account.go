package usecase

import (
	"context"
	"stoneBanking/app/application/validations"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
)

func (usecase *usecase) Create(ctx context.Context, accountData input.CreateAccountVO) (*output.AccountOutputVO, error) {
	var accountOutput output.AccountOutputVO
	var err error
	var tempAccount = &account.Account{}

	accountData, err = validations.ValidateAccountInput(accountData)

	if err != nil {
		return &accountOutput, err
	}

	tempAccount, err = usecase.accountRepository.GetByCPF(ctx, accountData.CPF, tempAccount)

	if err == nil {
		accountOutput = output.AccountToOutput(*tempAccount)
		return &accountOutput, errorAccountCPFExists
	}

	account := input.GenerateAccount(accountData)

	accountResult, err := usecase.accountRepository.Create(ctx, &account)

	if err != nil {
		return &accountOutput, err
	}

	accountOutput = output.AccountToOutput(*accountResult)

	return &accountOutput, err
}
