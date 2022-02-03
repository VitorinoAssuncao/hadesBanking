package account

import (
	"context"
	"errors"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"

	"github.com/jackc/pgx/v4"
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
		ctx,
		sqlQuery,
		accountExternalID.ToUUID(),
	)
	err := result.Scan(&newAccount.ID, &newAccount.ExternalID, &newAccount.Name, &newAccount.CPF, &newAccount.Secret, &newAccount.Balance, &newAccount.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return account.Account{}, customError.ErrorAccountIDNotFound
		}

		return account.Account{}, err
	}

	return newAccount, nil
}
