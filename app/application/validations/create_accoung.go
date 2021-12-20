package validations

import (
	"regexp"
	"stoneBanking/app/application/vo/input"
)

func ValidateAccountInput(accountData input.CreateAccountVO) (input.CreateAccountVO, error) {
	if !nameIsNotEmpty(accountData.Name) {
		return accountData, errorAccountNameRequired
	}

	if !cpfIsNotEmpty(accountData.CPF) {
		return accountData, errorAccountCPFRequired
	}

	if !cpfIsJustNumbers(accountData.CPF) {
		return accountData, errorAccountCPFNotNumbers
	}

	if !cpfIsValid(accountData.CPF) {
		return accountData, errorAccountCPFInvalid
	}

	if !secretIsNotEmpty(accountData.Secret) {
		return accountData, errorAccountSecretRequired
	}

	if !balanceIsPositive(accountData.Balance) {
		return accountData, errorAccountBalanceInvalid
	}

	return accountData, nil
}

func nameIsNotEmpty(name string) bool {
	return name != ""
}

func cpfIsNotEmpty(cpf string) bool {
	return cpf != ""
}

func cpfIsJustNumbers(cpf string) bool {
	p, _ := regexp.Compile("[^0-9]+")
	return !(p.Match([]byte(cpf)))
}

func cpfIsValid(cpf string) bool {
	//TODO implementar regra de validação de cpf
	return true
}

func secretIsNotEmpty(secret string) bool {
	return secret != ""
}

func balanceIsPositive(balance int) bool {
	return balance >= 0
}
