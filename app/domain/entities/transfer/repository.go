package transfer

import (
	"context"
	"stoneBanking/app/domain/types"
)

type Repository interface {
	Create(ctx context.Context, transfer Transfer) (Transfer, error)
	GetAllByAccountID(ctx context.Context, accountID types.InternalID) ([]Transfer, error)
}
