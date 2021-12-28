package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
)

func (repository accountRepository) GetByID(ctx context.Context, accountID types.AccountID) (account.Account, error) {

	var sqlQuery = `
	SELECT 
		id, name, cpf, secret, balance, created_at
	FROM
		accounts
	WHERE
			id = $1
	`
	var newAccount = account.Account{}

	result := repository.db.QueryRow(
		sqlQuery,
		accountID,
	)
	err := result.Scan(&newAccount.ID, &newAccount.Name, &newAccount.CPF, &newAccount.Secret, &newAccount.Balance, &newAccount.CreatedAt)

	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
