package account

import (
	"context"
	"database/sql"
	"errors"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func (repository accountRepository) GetCredentialByCPF(ctx context.Context, accountCPF string) (account.Account, error) {

	const sqlQuery = `
	SELECT 
		external_id,cpf, secret
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

	err := result.Scan(&newAccount.ExternalID, &newAccount.CPF, &newAccount.Secret)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return account.Account{}, customError.ErrorAccountCPFNotFound
		}

		return account.Account{}, err
	}

	return newAccount, nil
}
