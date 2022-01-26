package output

import (
	"stoneBanking/app/domain/entities/account"
)

type AccountOutputVO struct {
	ID         string  `json:"id" example:"123"`
	Name       string  `json:"name" example:"João da Silva"`
	CPF        string  `json:"cpf" example:"600.246.058-67"`
	Balance    float64 `json:"balance" example:"10.00"`
	Created_At string  `json:"created_at" example:"12/05/2021 00:01:01" `
}

type AccountBalanceVO struct {
	Balance float64 `json:"balance" example:"12.34"`
}

func ToOutput(account account.Account) AccountOutputVO {
	accountOutput := AccountOutputVO{
		ID:         string(account.ExternalID),
		Name:       account.Name,
		CPF:        account.CPF.ToString(),
		Balance:    account.Balance.ToFloat(),
		Created_At: account.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
	return accountOutput
}

func ToOutputs(accounts []account.Account) []AccountOutputVO {
	var accountsOutput = make([]AccountOutputVO, 0)
	for _, account := range accounts {
		tempAccount := ToOutput(account)
		accountsOutput = append(accountsOutput, tempAccount)
	}
	return accountsOutput
}
