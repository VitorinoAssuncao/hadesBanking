package accounts

import (
	account_usecase "stoneBanking/app/application/usecase/account"
)

type Controller struct {
	usecase    account_usecase.Usecase
	signingKey string
}

func New(useCase account_usecase.Usecase, key string) Controller {
	return Controller{
		usecase:    useCase,
		signingKey: key,
	}
}
