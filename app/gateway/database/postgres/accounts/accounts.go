package account

import (
	"database/sql"

	"stoneBanking/app/domain/entities/account"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(database *sql.DB) account.Repository {
	return &accountRepository{
		db: database,
	}
}
