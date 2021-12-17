package account

import (
	"context"
	"hades/adapters/vo/input"
	"stoneBanking/application/entities/account"
	"stoneBanking/domain/vo/output"
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
