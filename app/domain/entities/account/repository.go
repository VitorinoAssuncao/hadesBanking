package account

import (
	"context"
	"stoneBanking/app/domain/types"
)

type Repository interface {
	Create(ctx context.Context, account Account) (Account, error)
	GetByID(ctx context.Context, accountID types.ExternalID) (Account, error)
	GetByCPF(ctx context.Context, accountCPF string) (Account, error)
	GetAll(ctx context.Context) ([]Account, error)
	UpdateBalance(ctx context.Context, value int, externalID types.ExternalID) error
}
