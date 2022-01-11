package server

import (
	"stoneBanking/app/application/usecase/account"
	"stoneBanking/app/application/usecase/transfer"
	"stoneBanking/app/common/utils/config"
)

type UseCaseWrapper struct {
	Accounts account.Usecase
	Transfer transfer.Usecase
}

func NewUseCaseWrapper(wrapper *RepositorieWrapper, cfg config.Config) *UseCaseWrapper {
	return &UseCaseWrapper{
		Accounts: account.New(wrapper.Account, cfg.SigningKey),
		Transfer: transfer.New(wrapper.Transfer, wrapper.Account),
	}
}
