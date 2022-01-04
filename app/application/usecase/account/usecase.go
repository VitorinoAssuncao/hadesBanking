package account

import (
	"context"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
)

type Usecase interface {
	Create(ctx context.Context, accountData input.CreateAccountVO) (*output.AccountOutputVO, error)
	GetBalance(ctx context.Context, accountID string) (output.AccountBalanceVO, error)
	GetAll(ctx context.Context) ([]output.AccountOutputVO, error)
	LoginUser(ctx context.Context, loginInput input.LoginVO) (output.LoginOutputVO, error)
}

type usecase struct {
	accountRepository account.Repository
}

func New(account account.Repository) *usecase {
	return &usecase{
		accountRepository: account,
	}
}
