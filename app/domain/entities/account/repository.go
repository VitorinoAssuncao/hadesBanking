package account

import (
	"context"
	"stoneBanking/app/domain/types"
)

type Repository interface {
	Create(ctx context.Context, account *Account) (*Account, error)
	GetByID(ctx context.Context, accountID types.AccountID) (Account, error)
	GetByCPF(ctx context.Context, accountCPF string) (Account, error)
	GetAll(ctx context.Context) ([]Account, error)
}
