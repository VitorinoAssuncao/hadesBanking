package validations

import (
	"regexp"
	"stoneBanking/app/application/vo/input"
)

func ValidateAccountInput(accountData input.CreateAccountVO) (input.CreateAccountVO, error) {
	return accountData, nil
}

func nameIsNotEmpty(name string) bool {
	if name == "" {
		return false
	}
	return true
}

func cpfIsNotEmpty(cpf string) bool {
	if cpf == "" {
		return false
	}
	return true
}

func cpfIsJustNumbers(cpf string) bool {
	p, _ := regexp.Compile("[^0-9]+")
	if p.Match([]byte(cpf)) {
		return false
	}
	return true
}

func cpfIsValid(cpf string) bool {
	//TODO implementar regra de validação de cpf
	return true
}

func secretIsNotEmpty(secret string) bool {
	if secret == "" {
		return false
	}
	return true
}

func balanceIsPositive(balance int) bool {
	if balance < 0 {
		return false
	}
	return true
}
