package input

import (
	"stoneBanking/app/common/utils"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"time"
)

type CreateAccountVO struct {
	Name    string `json:"name" example:"João da Silva"`
	CPF     string `json:"cpf" example:"600.246.058-67"`
	Secret  string `json:"secret" example:"123456"`
	Balance int    `json:"balance" example:"1000"`
}

func (inputAccount CreateAccountVO) GenerateAccount() account.Account {
	account := account.Account{
		Name:      inputAccount.Name,
		CPF:       utils.TrimCPF(inputAccount.CPF),
		Secret:    utils.HashPassword(inputAccount.Secret),
		Balance:   types.Money(inputAccount.Balance),
		CreatedAt: time.Now(),
	}
	return account
}
