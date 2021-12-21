package server

import (
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	postgres_account "stoneBanking/app/gateway/database/postgres/accounts"
)

type RepositorieWrapper struct {
	Account account.Repository
}

func NewPostgresRepositoryWrapper(db *sql.DB) *RepositorieWrapper {
	return &RepositorieWrapper{
		Account: postgres_account.NewAccountRepository(db),
	}
}
