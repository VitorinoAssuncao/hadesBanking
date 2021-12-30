package transfer

import (
	"context"
	"stoneBanking/app/domain/types"
)

type Repository interface {
	Create(ctx context.Context, transfer Transfer) (Transfer, error)
	GetByID(ctx context.Context, transferID types.TransferExternalID) (Transfer, error)
	GetAll(ctx context.Context) ([]Transfer, error)
	GetAllByAccountID(ctx context.Context, accountID types.TransferAccountID) ([]Transfer, error)
}
