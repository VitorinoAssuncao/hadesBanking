package transfer

import (
	"stoneBanking/app/application/usecase/transfer"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase   transfer.Usecase
	tokenRepo token.Repository
}

func New(useCase transfer.Usecase, token token.Repository) Controller {
	return Controller{
		usecase:   useCase,
		tokenRepo: token,
	}
}
