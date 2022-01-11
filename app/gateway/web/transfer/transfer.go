package transfer

import "stoneBanking/app/application/usecase/transfer"

type Controller struct {
	usecase    transfer.Usecase
	signingKey string
}

func New(u transfer.Usecase, key string) Controller {
	return Controller{
		usecase:    u,
		signingKey: key,
	}
}
