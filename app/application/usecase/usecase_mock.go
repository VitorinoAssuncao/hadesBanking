package usecase

import (
	"stoneBanking/app/domain/entities/account"
	"time"
)

type MockUseCase struct {
	MockAccount  account.Account
	MockAccounts []account.Account
	Err          error
}

func (mock MockUseCase) SetupMockAccountBalance() MockUseCase {
	mock.MockAccount = account.Account{
		ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
		Name:       "João da Silva",
		Secret:     "12345",
		Balance:    1235,
		Created_at: time.Now(),
	}
	return mock
}
