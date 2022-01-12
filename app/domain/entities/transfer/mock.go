package transfer

import (
	"context"
	"stoneBanking/app/domain/types"
)

type RepositoryMock struct {
	CreateFunc            func(ctx context.Context, transfer Transfer) (Transfer, error)
	GetByIDFunc           func(ctx context.Context, transferID types.ExternalID) (Transfer, error)
	GetAllFunc            func(ctx context.Context) ([]Transfer, error)
	GetAllByAccountIDFunc func(ctx context.Context, accountID types.InternalID) ([]Transfer, error)
}

func (r *RepositoryMock) Create(ctx context.Context, transfer Transfer) (Transfer, error) {
	return r.CreateFunc(ctx, transfer)
}

func (r *RepositoryMock) GetByID(ctx context.Context, transferID types.ExternalID) (Transfer, error) {
	return r.GetByIDFunc(ctx, transferID)
}

func (r *RepositoryMock) GetAll(ctx context.Context) ([]Transfer, error) {
	return r.GetAllFunc(ctx)
}

func (r *RepositoryMock) GetAllByAccountID(ctx context.Context, accountID types.InternalID) ([]Transfer, error) {
	return r.GetAllByAccountIDFunc(ctx, accountID)
}
