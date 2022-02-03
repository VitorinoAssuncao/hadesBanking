package transfer

import (
	"stoneBanking/app/domain/entities/transfer"

	"github.com/jackc/pgx/v4"
)

type transferRepository struct {
	db *pgx.Conn
}

func NewTransferRepository(database *pgx.Conn) transfer.Repository {
	return &transferRepository{
		db: database,
	}
}
