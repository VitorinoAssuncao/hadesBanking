package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) GetByCPF(ctx context.Context, accountCPF string) (account.Account, error) {

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
	var newAccount = account.Account{}
	err := result.Scan(&newAccount.ID, &newAccount.Name, &newAccount.Cpf, &newAccount.Secret, &newAccount.Balance, &newAccount.Created_at)

	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
