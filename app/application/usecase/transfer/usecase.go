package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

type Usecase interface {
	Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error)
	GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]transfer.Transfer, error)
}

type usecase struct {
	transferRepository transfer.Repository
	accountRepository  account.Repository
	logRepository      logHelper.Repository
}

func New(transfer transfer.Repository, account account.Repository, log logHelper.Repository) *usecase {
	return &usecase{
		transferRepository: transfer,
		accountRepository:  account,
		logRepository:      log,
	}
}
