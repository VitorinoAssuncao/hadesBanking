package accounts

import (
	account_usecase "stoneBanking/app/application/usecase/account"
)

type Controller struct {
	usecase account_usecase.Usecase
}

func New(useCase account_usecase.Usecase) Controller {
	return Controller{
		usecase: useCase,
	}
}
