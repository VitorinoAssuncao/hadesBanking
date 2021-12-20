package input

import (
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"time"

	"github.com/google/uuid"
)

type CreateAccountVO struct {
	Name    string `json:"name" example:"João da Silva"`
	CPF     string `json:"cpf" example:"600.246.058-67"`
	Secret  string `json:"secret" example:"123456"`
	Balance int    `json:"balance" example:"1000"`
}

func GenerateAccount(inputAccount CreateAccountVO) account.Account {
	tempID, _ := uuid.NewRandom()
	account := account.Account{
		ID:         types.AccountID(tempID.String()),
		Name:       inputAccount.Name,
		Cpf:        inputAccount.CPF,
		Secret:     inputAccount.Secret,
		Balance:    types.Money(inputAccount.Balance),
		Created_at: time.Now(),
	}
	return account
}
