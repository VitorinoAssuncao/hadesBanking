package transfer

import (
	"stoneBanking/app/domain/entities/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type transferRepository struct {
	db *pgxpool.Pool
}

func NewTransferRepository(database *pgxpool.Pool) transfer.Repository {
	return &transferRepository{
		db: database,
	}
}
