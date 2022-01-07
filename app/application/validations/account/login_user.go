package validations

import (
	"stoneBanking/app/application/vo/input"
)

func ValidateLoginInputData(input input.LoginVO) error {
	if input.CPF == "" {
		return ErrorAccountCPFRequired
	}

	if input.Secret == "" {
		return ErrorAccountSecretRequired
	}

	return nil
}
