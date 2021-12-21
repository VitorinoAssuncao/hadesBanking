package accounts

import (
	"stoneBanking/app/application/usecase"
)

type Controller struct {
	usecase usecase.Usecase
}

func New(useCase usecase.Usecase) Controller {
	return Controller{
		usecase: useCase,
	}
}
