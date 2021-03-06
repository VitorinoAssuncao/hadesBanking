package input

import (
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/account/vo/input"
)

func ValidateAccountInput(accountData input.CreateAccountVO) []error {
	_, err := validateName(accountData.Name)
	var errs = make([]error, 0)
	if err != nil {
		errs = append(errs, err)
	}

	_, err = validateCPF(accountData.CPF.ToString())
	if err != nil {
		errs = append(errs, err)
	}

	_, err = validateSecret(accountData.Secret.ToString())
	if err != nil {
		errs = append(errs, err)
	}

	_, err = validateBalance(accountData.Balance)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func validateName(name string) (bool, error) {
	if name == "" {
		return false, customError.ErrorAccountNameRequired
	}

	return true, nil
}

func validateCPF(cpf string) (bool, error) {
	if cpf == "" {
		return false, customError.ErrorAccountCPFRequired
	}

	if len(cpf) != 11 {
		return false, customError.ErrorAccountCPFWrongSize
	}

	if !cpfIsValid(cpf) {
		return false, customError.ErrorAccountCPFInvalid
	}

	if !cpfIsNotATestValue(cpf) {
		return false, customError.ErrorAccountCPFTestNumber
	}
	return true, nil
}

func validateSecret(secret string) (bool, error) {
	if secret == "" {
		return false, customError.ErrorAccountSecretRequired
	}

	return true, nil
}

func validateBalance(balance int) (bool, error) {
	if balance < 0 {
		return false, customError.ErrorAccountBalanceInvalid
	}

	return true, nil
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

func calculateFirstVerifyingDigit(values []int) int {
	var total int
	for index, value := range values {
		total += ((index + 1) * value)
	}
	const divisionFactor = 11
	result := total % divisionFactor

	const decimalConversion = 10
	if result == decimalConversion {
		return 0
	}

	return result
}

func calculateSecondVerifyingDigit(values []int) int {
	var total int
	for index, value := range values {
		total += ((index) * value)
	}
	const divisionFactor = 11
	result := total % divisionFactor

	const decimalConversion = 10
	if result == decimalConversion {
		return 0
	}

	return result
}

func cpfIsNotATestValue(cpf string) bool {
	// Validação leva em conta se CPF apresenta dados inválidos de teste (todos os números iguais ou padrão sequencial 12345678901)
	const cpfTestExample = "12345678901"
	if cpf == cpfTestExample || cpf[0:5] == cpf[5:10] {
		return false
	}

	return true
}
