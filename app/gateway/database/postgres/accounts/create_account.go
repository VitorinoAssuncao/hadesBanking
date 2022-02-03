package account

import (
	"context"

	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) Create(ctx context.Context, newAccount account.Account) (account.Account, error) {
	const sqlQuery = `
	INSERT INTO
			accounts (name, cpf, secret, balance)
	VALUES
			($1, $2, $3, $4)
	RETURNING
			id, external_id, created_at
	`
	row := repository.db.QueryRow(
		ctx,
		sqlQuery,
		newAccount.Name,
		newAccount.CPF,
		newAccount.Secret,
		newAccount.Balance.ToInt(),
	)

	err := row.Scan(&newAccount.ID, &newAccount.ExternalID, &newAccount.CreatedAt)

	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
