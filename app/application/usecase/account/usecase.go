package account

import (
	"context"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
)

type UseCase interface {
	Create(ctx context.Context, accountData input.CreateAccountVO) (*output.AccountOutputVO, error)
	GetBalance(ctx context.Context, accountID string) (output.AccountBalanceVO, error)
	GetAll(ctx context.Context) ([]output.AccountOutputVO, error)
}

type usecase struct {
	account account.Repository
}

func New(account account.Repository) *usecase {
	return &usecase{
		account: account,
	}
}
