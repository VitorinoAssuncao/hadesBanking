package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
)

type Usecase interface {
	Create(ctx context.Context, accountData account.Account) (account.Account, error)
	GetBalance(ctx context.Context, accountID string) (float64, error)
	GetAll(ctx context.Context) ([]account.Account, error)
	LoginUser(ctx context.Context, loginInput account.Account) (string, error)
}

type usecase struct {
	accountRepository account.Repository
	tokenRepository   token.Repository
	logger            logHelper.Logger
}

func New(account account.Repository, token token.Repository, logger logHelper.Logger) *usecase {
	return &usecase{
		accountRepository: account,
		tokenRepository:   token,
		logger:            logger,
	}
}
