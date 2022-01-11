package accounts

import (
	account_usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase   account_usecase.Usecase
	tokenRepo token.TokenInterface
}

func New(useCase account_usecase.Usecase, tokenRepository token.TokenInterface) Controller {
	return Controller{
		usecase:   useCase,
		tokenRepo: tokenRepository,
	}
}
