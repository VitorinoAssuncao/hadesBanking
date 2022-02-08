package transfer

import (
	"context"
	"stoneBanking/app/domain/types"
	"sync/atomic"
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

type ParallelMock struct {
	CreateFunc            func(ctx context.Context, transfer Transfer) (Transfer, error)
	GetByIDFunc           func(ctx context.Context, transferID types.ExternalID) (Transfer, error)
	GetAllFunc            func(ctx context.Context) ([]Transfer, error)
	GetAllByAccountIDFunc func(ctx context.Context, accountID types.InternalID) ([]Transfer, error)
	Count                 int32
	WaitChan              chan bool
}

func (pr *ParallelMock) Create(ctx context.Context, transfer Transfer) (Transfer, error) {
	atomic.AddInt32(&pr.Count, 1)
	if <-pr.WaitChan {
		close(pr.WaitChan)
	}
	return pr.CreateFunc(ctx, transfer)
}

func (pr *ParallelMock) GetByID(ctx context.Context, transferID types.ExternalID) (Transfer, error) {
	return pr.GetByIDFunc(ctx, transferID)
}

func (pr *ParallelMock) GetAll(ctx context.Context) ([]Transfer, error) {
	return pr.GetAllFunc(ctx)
}

func (pr *ParallelMock) GetAllByAccountID(ctx context.Context, accountID types.InternalID) ([]Transfer, error) {
	return pr.GetAllByAccountIDFunc(ctx, accountID)
}
