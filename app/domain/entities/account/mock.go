package account

import (
	"context"
	"stoneBanking/app/domain/types"
)

type RepositoryMock struct {
	CreateFunc                func(ctx context.Context, account Account) (Account, error)
	GetByIDFunc               func(ctx context.Context, accountID types.ExternalID) (Account, error)
	GetByCPFFunc              func(ctx context.Context, accountCPF string) (Account, error)
	GetBalanceByAccountIDFunc func(ctx context.Context, accountID types.ExternalID) (types.Money, error)
	GetAllFunc                func(ctx context.Context) ([]Account, error)
	UpdateBalanceFunc         func(ctx context.Context, value int, externalID types.ExternalID) error
}

func (r *RepositoryMock) Create(ctx context.Context, account Account) (Account, error) {
	return r.CreateFunc(ctx, account)
}

func (r *RepositoryMock) GetByID(ctx context.Context, accountID types.ExternalID) (Account, error) {
	return r.GetByIDFunc(ctx, accountID)
}
func (r *RepositoryMock) GetByCPF(ctx context.Context, accountCPF string) (Account, error) {
	return r.GetByCPFFunc(ctx, accountCPF)
}

func (r *RepositoryMock) GetBalanceByAccountID(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
	return r.GetBalanceByAccountIDFunc(ctx, accountID)
}

func (r *RepositoryMock) GetAll(ctx context.Context) ([]Account, error) {
	return r.GetAllFunc(ctx)
}
func (r *RepositoryMock) UpdateBalance(ctx context.Context, value int, externalID types.ExternalID) error {
	return r.UpdateBalanceFunc(ctx, value, externalID)
}
