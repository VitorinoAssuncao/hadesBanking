package transfer

import (
	"stoneBanking/app/application/usecase/transfer"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase transfer.Usecase
	token   token.Authenticator
	log     logHelper.Logger
}

func New(useCase transfer.Usecase, token token.Authenticator, log logHelper.Logger) Controller {
	return Controller{
		usecase: useCase,
		token:   token,
		log:     log,
	}
}
