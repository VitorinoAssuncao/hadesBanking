package postgres_account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) Create(ctx context.Context, account account.Account) (account.Account, error) {
	var sqlQuery = `
	INSERT INTO
			accounts (id, name, cpf, secret, balance, created_at)
	VALUES
			($1, $2, $3, $4, $5, $6)
	`
	_, err := repository.db.Exec(
		sqlQuery,
		account.ID,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance.ToInt(),
		account.Created_at)

	if err != nil {
		return account, err
	}
	return account, nil
}
