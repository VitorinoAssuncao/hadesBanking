package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

type Usecase interface {
	Create(ctx context.Context, accountData account.Account) (account.Account, error)
	GetBalance(ctx context.Context, accountID string) (float64, error)
	GetAll(ctx context.Context) ([]account.Account, error)
	LoginUser(ctx context.Context, loginInput account.Account) (string, error)
}

type usecase struct {
	accountRepository account.Repository
}

func New(account account.Repository) *usecase {
	return &usecase{
		accountRepository: account,
	}
}
