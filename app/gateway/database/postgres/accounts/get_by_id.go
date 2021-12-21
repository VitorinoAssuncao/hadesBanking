package postgres_account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
)

func (repository accountRepository) GetByID(ctx context.Context, accountID types.AccountID, account *account.Account) (*account.Account, error) {

	var sqlQuery = `
	SELECT 
		id, name, cpf, secret, balance, created_at
	FROM
		accounts
	WHERE
			id = $1
	`
	result := repository.db.QueryRow(
		ctx,
		sqlQuery,
		accountID,
	)
	err := result.Scan(&account.ID, &account.Name, &account.Cpf, &account.Secret, &account.Balance, &account.Created_at)

	if err != nil {
		return nil, errorAccountIDNotFound
	}

	return account, nil
}
