package authentication

import (
	auth_usecase "stoneBanking/app/application/usecase/authentication"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase auth_usecase.Usecase
	token   token.Authenticator
	log     logHelper.Logger
}

func New(useCase auth_usecase.Usecase, token token.Authenticator, log logHelper.Logger) Controller {
	return Controller{
		usecase: useCase,
		token:   token,
		log:     log,
	}
}
