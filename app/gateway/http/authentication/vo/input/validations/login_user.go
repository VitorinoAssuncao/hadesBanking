package input

import (
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/authentication/vo/input"
)

func ValidateLoginInputData(input input.LoginVO) error {
	if input.CPF == "" {
		return customError.ErrorAccountLogin
	}

	if input.Secret == "" {
		return customError.ErrorAccountLogin
	}

	return nil
}
