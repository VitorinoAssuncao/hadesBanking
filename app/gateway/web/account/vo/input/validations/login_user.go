package input

import (
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/account/vo/input"
)

func ValidateLoginInputData(input input.CreateAccountVO) error {
	if input.CPF == "" {
		return customError.ErrorAccountLogin
	}

	if input.Secret == "" {
		return customError.ErrorAccountLogin
	}

	return nil
}
