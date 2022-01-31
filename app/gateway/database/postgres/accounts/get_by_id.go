package account

import (
	"context"

	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
)

func (repository accountRepository) GetByID(ctx context.Context, accountExternalID types.ExternalID) (account.Account, error) {

	const sqlQuery = `
	SELECT 
		id, external_id, name, cpf, secret, balance, created_at
	FROM
		accounts
	WHERE
		external_id = $1
	`
	var newAccount = account.Account{}

	result := repository.db.QueryRow(
		sqlQuery,
		accountExternalID,
	)
	err := result.Scan(&newAccount.ID, &newAccount.ExternalID, &newAccount.Name, &newAccount.CPF, &newAccount.Secret, &newAccount.Balance, &newAccount.CreatedAt)

	if err != nil {
		return account.Account{}, err
	}

	return newAccount, nil
}
