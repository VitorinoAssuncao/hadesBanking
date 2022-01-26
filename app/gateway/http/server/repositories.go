package server

import (
	"database/sql"
	commonLog "stoneBanking/app/common/utils/logger"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/entities/transfer"
	postgresAccount "stoneBanking/app/gateway/database/postgres/accounts"
	postgresTransfer "stoneBanking/app/gateway/database/postgres/transfers"
	webToken "stoneBanking/app/gateway/http/token"
)

type RepositorieWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
	Token    token.Authenticator
	Log      logHelper.Logger
}

func NewPostgresRepositoryWrapper(db *sql.DB, signKey string, log logHelper.Logger) *RepositorieWrapper {
	return &RepositorieWrapper{
		Account:  postgresAccount.NewAccountRepository(db),
		Transfer: postgresTransfer.NewTransferRepository(db),
		Token:    webToken.NewTokenAuthenticator(signKey),
		Log:      commonLog.NewLogger(),
	}
}