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
	webToken "stoneBanking/app/gateway/web/token"
)

type RepositorieWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
	Token    token.Repository
	Log      logHelper.Repository
}

func NewPostgresRepositoryWrapper(db *sql.DB, signKey string) *RepositorieWrapper {
	return &RepositorieWrapper{
		Account:  postgresAccount.NewAccountRepository(db),
		Transfer: postgresTransfer.NewTransferRepository(db),
		Token:    webToken.NewTokenRepository(signKey),
		Log:      commonLog.NewLogRepository(),
	}
}
