package transfer

import (
	"context"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

type Usecase interface {
	Create(ctx context.Context, transferData input.CreateTransferVO) (*output.TransferOutputVO, error)
	GetAllByAccountID(ctx context.Context, accountID types.ExternalID) ([]output.TransferOutputVO, error)
}

type usecase struct {
	transferRepository transfer.Repository
	accountRepository  account.Repository
}

func New(transfer transfer.Repository, account account.Repository) *usecase {
	return &usecase{
		transferRepository: transfer,
		accountRepository:  account,
	}
}
