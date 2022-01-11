package server

import (
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/entities/transfer"
	postgresAccount "stoneBanking/app/gateway/database/postgres/accounts"
	postgresTransfer "stoneBanking/app/gateway/database/postgres/transfers"
	webToken "stoneBanking/app/gateway/web/token"
)

type RepositorieWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
	Token    token.Repository
}

func NewPostgresRepositoryWrapper(db *sql.DB, signKey string) *RepositorieWrapper {
	return &RepositorieWrapper{
		Account:  postgresAccount.NewAccountRepository(db),
		Transfer: postgresTransfer.NewTransferRepository(db),
		Token:    webToken.NewTokenRepository(signKey),
	}
}
