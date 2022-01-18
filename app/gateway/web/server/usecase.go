package server

import (
	"stoneBanking/app/application/usecase/account"
	"stoneBanking/app/application/usecase/transfer"
)

type UseCaseWrapper struct {
	Accounts account.Usecase
	Transfer transfer.Usecase
}

func NewUseCaseWrapper(wrapper *RepositorieWrapper) *UseCaseWrapper {
	return &UseCaseWrapper{
		Accounts: account.New(wrapper.Account, wrapper.Token, wrapper.Log),
		Transfer: transfer.New(wrapper.Transfer, wrapper.Account),
	}
}
