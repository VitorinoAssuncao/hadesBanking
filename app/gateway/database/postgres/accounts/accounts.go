package account

import (
	"stoneBanking/app/domain/entities/account"

	"github.com/jackc/pgx/v4"
)

type accountRepository struct {
	db *pgx.Conn
}

func NewAccountRepository(database *pgx.Conn) account.Repository {
	return &accountRepository{
		db: database,
	}
}
