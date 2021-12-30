package validations

import (
	"stoneBanking/app/application/vo/input"
)

func ValidateAccountInput(accountData input.CreateAccountVO) (input.CreateAccountVO, error) {
	if !nameIsNotEmpty(accountData.Name) {
		return accountData, errorAccountNameRequired
	}

	if !cpfIsRightSIze(accountData.CPF) {
		return accountData, errorAccountCPFWrongSize
	}

	if !cpfIsNotEmpty(accountData.CPF) {
		return accountData, errorAccountCPFRequired
	}

	if cpfIsNotATestValue(accountData.CPF) {
		return accountData, errorAccountCPFInvalid
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

func cpfIsRightSIze(cpf string) bool {
	return len(cpf) == 11
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

func cpfIsNotATestValue(cpf string) bool {
	// Validação leva em conta se CPF apresenta dados inválidos de teste (todos os números iguais ou padrão sequencial 12345678901)
	if cpf == "12345678901" {
		return false
	}

	if cpf[0:5] == cpf[5:10] {
		return false
	}

	return true
}
