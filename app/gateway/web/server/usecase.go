package server

import (
	"stoneBanking/app/application/usecase"
)

type UseCaseWrapper struct {
	Accounts usecase.Usecase
}

func NewUseCaseWrapper(wrapper *RepositorieWrapper) *UseCaseWrapper {
	return &UseCaseWrapper{
		Accounts: usecase.New(wrapper.Account),
	}
}
