package input

import (
	"stoneBanking/app/common/utils"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateAccountVO struct {
	Name    string `json:"name" example:"João da Silva"`
	CPF     string `json:"cpf" example:"600.246.058-67"`
	Secret  string `json:"secret" example:"123456"`
	Balance int    `json:"balance" example:"1000"`
}

func GenerateAccount(inputAccount CreateAccountVO) account.Account {
	account := account.Account{
		Name:      inputAccount.Name,
		CPF:       utils.TrimCPF(inputAccount.CPF),
		Secret:    hashPassword(inputAccount.Secret),
		Balance:   types.Money(inputAccount.Balance),
		CreatedAt: time.Now(),
	}
	return account
}

func hashPassword(text string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(bytes)
}

func ValidateHash(accountSecret, loginSecret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(accountSecret), []byte(loginSecret))
	return err == nil
}
