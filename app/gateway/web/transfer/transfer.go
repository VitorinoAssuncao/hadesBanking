package transfer

import (
	"stoneBanking/app/application/usecase/transfer"
	"stoneBanking/app/domain/entities/token"
)

type Controller struct {
	usecase   transfer.Usecase
	tokenRepo token.TokenInterface
}

func New(u transfer.Usecase, tokenRepository token.TokenInterface) Controller {
	return Controller{
		usecase:   u,
		tokenRepo: tokenRepository,
	}
}
