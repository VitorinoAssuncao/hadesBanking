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
	cpfArray := make([]int, 0)

	for _, i := range cpf {
		value := int(i) - '0' //subtração devolve o valor do inteiro corretamente (49 - valor de 1, - valor do 0 (48))
		cpfArray = append(cpfArray, value)
	}

	firstDigit := calculateFirstVerifyingDigit(cpfArray[0:9])
	secondDigit := calculateSecondVerifyingDigit(cpfArray[0:10])
	return firstDigit == cpfArray[9] && secondDigit == cpfArray[10]
}

func secretIsNotEmpty(secret string) bool {
	return secret != ""
}

func balanceIsPositive(balance int) bool {
	return balance >= 0
}

func calculateFirstVerifyingDigit(values []int) int {
	var total int
	for index, value := range values {
		total += ((index + 1) * value)
	}
	result := total % 11

	if result == 10 {
		return 0
	}

	return result
}

func calculateSecondVerifyingDigit(values []int) int {
	var total int
	for index, value := range values {
		total += ((index) * value)
	}
	result := total % 11

	if result == 10 {
		return 0
	}

	return result
}
