package output

import (
	"hades/domain/entities"
)

type AccountOutputVO struct {
	ID         string  `json:"account_id" example:"123"`
	Name       string  `json:"name" example:"João da Silva"`
	CPF        string  `json:"cpf" example:"600.246.058-67"`
	Balance    float64 `json:"balance" example:"10.00"`
	Created_At string  `json:"created_at" example:"12/05/2021 00:01:01" `
}

type AccountBalanceVO struct {
	Balance float64 `json:"balance" example:"12.34"`
}

func AccountToOutput(account entities.Account) AccountOutputVO {
	accountOutput := AccountOutputVO{
		ID:         string(account.ID),
		Name:       account.Name,
		CPF:        account.Cpf,
		Balance:    account.Balance.ToFloat(),
		Created_At: account.Created_at.Format("2006-01-02T15:04:05Z07:00"),
	}
	return accountOutput
}
