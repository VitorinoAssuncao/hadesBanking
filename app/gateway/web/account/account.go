package accounts

import (
	"stoneBanking/app/application/usecase/account"
)

type Controller struct {
	usecase account.Usecase
}

func New(useCase account.Usecase) Controller {
	return Controller{
		usecase: useCase,
	}
}
