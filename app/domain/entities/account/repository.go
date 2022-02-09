package account

import (
	"context"
	"stoneBanking/app/domain/types"
)

type Repository interface {
	Create(ctx context.Context, account Account) (Account, error)
	GetByID(ctx context.Context, accountID types.ExternalID) (Account, error)
	GetCredentialByCPF(ctx context.Context, accountCPF string) (Account, error)
	GetBalanceByAccountID(ctx context.Context, accountID types.ExternalID) (types.Money, error)
	GetAll(ctx context.Context) ([]Account, error)
}
