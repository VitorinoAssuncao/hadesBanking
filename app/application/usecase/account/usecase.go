package account

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/types"
)

type Usecase interface {
	Create(ctx context.Context, accountData account.Account) (account.Account, error)
	GetBalance(ctx context.Context, accountID string) (types.Money, error)
	GetAll(ctx context.Context) ([]account.Account, error)
}

type usecase struct {
	accountRepository account.Repository
	token             token.Authenticator
	logger            logHelper.Logger
}

func New(account account.Repository, token token.Authenticator, logger logHelper.Logger) *usecase {
	return &usecase{
		accountRepository: account,
		token:             token,
		logger:            logger,
	}
}
