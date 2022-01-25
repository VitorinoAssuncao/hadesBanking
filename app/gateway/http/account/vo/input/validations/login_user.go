package input

import (
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/account/vo/input"
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
