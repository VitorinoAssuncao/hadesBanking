package account

import (
	account_usecase "stoneBanking/app/application/usecase/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase account_usecase.Usecase
	token   token.Authenticator
	log     logHelper.Logger
}

func New(useCase account_usecase.Usecase, token token.Authenticator, log logHelper.Logger) Controller {
	return Controller{
		usecase: useCase,
		token:   token,
		log:     log,
	}
}
