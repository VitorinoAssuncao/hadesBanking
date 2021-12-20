package server

import "stoneBanking/app/application/usecase/account"

type UseCaseWrapper struct {
	Accounts account.Usecase
}

func NewUseCaseWrapper(wrapper *RepositorieWrapper) *UseCaseWrapper {
	return &UseCaseWrapper{
		Accounts: account.New(wrapper.Account),
	}
}
