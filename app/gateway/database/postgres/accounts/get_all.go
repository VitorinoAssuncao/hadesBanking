package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) GetAll(ctx context.Context) ([]account.Account, error) {
	var tempAccount account.Account
	var accounts = make([]account.Account, 0)
	var sqlQuery = `
	SELECT 
		id, name, cpf, secret, balance, created_at
	FROM
		accounts
	`
	result, err := repository.db.Query(sqlQuery)
	if err != nil {
		return accounts, err
	}

	for result.Next() {
		err = result.Scan(&tempAccount.ID, &tempAccount.Name, &tempAccount.CPF, &tempAccount.Secret, &tempAccount.Balance, &tempAccount.CreatedAt)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, tempAccount)
	}

	return accounts, nil
}
