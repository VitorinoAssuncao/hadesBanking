package validations

import (
	"errors"
	"stoneBanking/app/application/vo/input"
)

func ValidateLoginInputData(input input.LoginVO) (result bool, err error) {
	if input.CPF == "" {
		return false, errors.New("cpf é um campo obrigatório")
	}

	if input.Secret == "" {
		return false, errors.New("senha é um campo obrigatório")
	}

	return true, nil
}
