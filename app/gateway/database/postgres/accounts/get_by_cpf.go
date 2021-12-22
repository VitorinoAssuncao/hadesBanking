package postgres_account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) GetByCPF(ctx context.Context, accountCPF string, account *account.Account) (*account.Account, error) {

	var sqlQuery = `
	SELECT 
		id, name, cpf, secret, balance, created_at
	FROM
		accounts
	WHERE
			cpf = $1
	`
	result := repository.db.QueryRow(
		sqlQuery,
		accountCPF,
	)

	err := result.Scan(&account.ID, &account.Name, &account.Cpf, &account.Secret, &account.Balance, &account.Created_at)

	if err != nil {
		return nil, err
	}

	return account, nil
}
