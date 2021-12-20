package server

import (
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/gateway/database/postgres"

	"github.com/jackc/pgx/v4"
)

type RepositorieWrapper struct {
	Account account.Repository
}

func NewPostgresRepositoryWrapper(db *pgx.Conn) *RepositorieWrapper {
	return &RepositorieWrapper{
		Account: postgres.NewAccountRepository(db),
	}
}
