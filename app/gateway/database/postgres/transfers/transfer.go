package transfer

import (
	"database/sql"
	"stoneBanking/app/domain/entities/transfer"
)

type transferRepository struct {
	db *sql.DB
}

func NewTransferRepository(database *sql.DB) transfer.Repository {
	return &transferRepository{
		db: database,
	}
}
