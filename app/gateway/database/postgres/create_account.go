package postgres

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"

	"github.com/jackc/pgx/v4"
)

type accountRepository struct {
	db *pgx.Conn
}

func (repository accountRepository) Create(ctx context.Context, account *account.Account) (*account.Account, error) {
	var sqlQuery = `
	INSERT INTO
			accounts (id, name, cpf, secret, balance, created_at)
	VALUES
			($1, $2, $3, $4, $5, $6)
	`
	_, err := repository.db.Exec(
		ctx,
		sqlQuery,
		account.ID,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance.ToInt(),
		account.Created_at)

	if err != nil {
		return nil, errorCreateAccount
	}
	return account, nil
}

func (r accountRepository) GetBalance(ctx context.Context, accountID types.AccountID) (types.Money, error) {
	return 0, nil
}

func (r accountRepository) GetAll(ctx context.Context) ([]account.Account, error) {
	return nil, nil
}

func (r accountRepository) GetByID(ctx context.Context, accountID types.AccountID) (*account.Account, error) {
	return nil, nil
}

func NewAccountRepository(connection *pgx.Conn) account.Repository {
	return &accountRepository{
		db: connection,
	}
}
