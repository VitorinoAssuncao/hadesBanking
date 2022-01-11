package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/account"
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
	signingKey         string
}

func New(transfer transfer.Repository, account account.Repository) *usecase {
	return &usecase{
		transferRepository: transfer,
		accountRepository:  account,
	}
}
