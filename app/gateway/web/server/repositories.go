package server

import (
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/entities/transfer"
	postgresAccount "stoneBanking/app/gateway/database/postgres/accounts"
	postgresTransfer "stoneBanking/app/gateway/database/postgres/transfers"
)

type RepositorieWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
	Token    token.TokenInterface
}

func NewPostgresRepositoryWrapper(db *sql.DB, tokenRepository token.TokenInterface) *RepositorieWrapper {
	return &RepositorieWrapper{
		Account:  postgresAccount.NewAccountRepository(db),
		Transfer: postgresTransfer.NewTransferRepository(db),
		Token:    tokenRepository,
	}
}
