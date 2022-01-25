package input

import (
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"time"
)

type CreateAccountVO struct {
	Name    string         `json:"name" example:"João da Silva"`
	CPF     types.Document `json:"cpf" example:"600.246.058-67"`
	Secret  types.Password `json:"secret" example:"123456"`
	Balance int            `json:"balance" example:"1000"`
}

func (inputAccount CreateAccountVO) GenerateAccount() account.Account {
	account := account.Account{
		Name:      inputAccount.Name,
		CPF:       types.Document(inputAccount.CPF.TrimCPF()),
		Secret:    inputAccount.Secret.Hash(),
		Balance:   types.Money(inputAccount.Balance),
		CreatedAt: time.Now(),
	}
	return account
}
