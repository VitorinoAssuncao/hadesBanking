package account

import (
	account_usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase   account_usecase.Usecase
	tokenRepo token.Repository
}

func New(useCase account_usecase.Usecase, token token.Repository) Controller {
	return Controller{
		usecase:   useCase,
		tokenRepo: token,
	}
}
