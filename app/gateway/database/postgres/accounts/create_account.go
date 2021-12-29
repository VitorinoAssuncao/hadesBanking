package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) Create(ctx context.Context, newAccount account.Account) (account.Account, error) {
	var sqlQuery = `
	INSERT INTO
			accounts (name, cpf, secret, balance, created_at)
	VALUES
			($1, $2, $3, $4, $5)
	RETURNING
			id, external_id
	`
	row := repository.db.QueryRow(
		sqlQuery,
		newAccount.Name,
		newAccount.CPF,
		newAccount.Secret,
		newAccount.Balance.ToInt(),
		newAccount.CreatedAt)

	err := row.Scan(&newAccount.ID, &newAccount.ExternalID)

	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
