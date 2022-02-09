package server

import (
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/entities/transfer"
	postgresAccount "stoneBanking/app/gateway/database/postgres/accounts"
	postgresTransfer "stoneBanking/app/gateway/database/postgres/transfers"
	webToken "stoneBanking/app/gateway/http/token"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
	Token    token.Authenticator
	Log      logHelper.Logger
}

func NewPostgresRepositoryWrapper(db *pgxpool.Pool, signKey string, log logHelper.Logger) *RepositoryWrapper {
	return &RepositoryWrapper{
		Account:  postgresAccount.NewAccountRepository(db),
		Transfer: postgresTransfer.NewTransferRepository(db),
		Token:    webToken.NewTokenAuthenticator(signKey),
		Log:      log,
	}
}
