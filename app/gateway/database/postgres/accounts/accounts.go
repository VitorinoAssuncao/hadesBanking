package account

import (
	"stoneBanking/app/domain/entities/account"

	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(database *pgxpool.Pool) account.Repository {
	return &accountRepository{
		db: database,
	}
}
