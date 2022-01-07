package transfer

import "stoneBanking/app/application/usecase/transfer"

type Controller struct {
	usecase transfer.Usecase
}

func New(u transfer.Usecase) Controller {
	return Controller{
		usecase: u,
	}
}
