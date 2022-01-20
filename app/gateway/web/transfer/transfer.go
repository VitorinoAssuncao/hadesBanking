package transfer

import (
	"stoneBanking/app/application/usecase/transfer"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase   transfer.Usecase
	tokenRepo token.Repository
	log       logHelper.Repository
}

func New(useCase transfer.Usecase, token token.Repository, log logHelper.Repository) Controller {
	return Controller{
		usecase:   useCase,
		tokenRepo: token,
		log:       log,
	}
}
