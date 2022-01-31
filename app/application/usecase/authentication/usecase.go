package authentication

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
)

type Usecase interface {
	LoginUser(ctx context.Context, loginInput account.Account) (string, error)
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
