package account

import (
	"context"

	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) GetAll(ctx context.Context) ([]account.Account, error) {
	const sqlQuery = `
	SELECT 
		id, external_id, name, cpf, secret, balance, created_at
	FROM
		accounts
	`
	var accounts = make([]account.Account, 0)

	result, err := repository.db.Query(ctx, sqlQuery)
	if err != nil {
		return accounts, err
	}

	var tempAccount account.Account

	for result.Next() {
		err = result.Scan(&tempAccount.ID, &tempAccount.ExternalID, &tempAccount.Name, &tempAccount.CPF, &tempAccount.Secret, &tempAccount.Balance, &tempAccount.CreatedAt)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, tempAccount)
	}

	return accounts, nil
}
