package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
)

func (repository accountRepository) Create(ctx context.Context, newAccount account.Account) (account.Account, error) {
	var sqlQuery = `
	INSERT INTO
			accounts (id, name, cpf, secret, balance)
	VALUES
			($1, $2, $3, $4, $5)
	`
	_, err := repository.db.Exec(
		sqlQuery,
		newAccount.ID,
		newAccount.Name,
		newAccount.CPF,
		newAccount.Secret,
		newAccount.Balance.ToInt(),
	)

	if err != nil {
		return account.Account{}, err
	}
	return newAccount, nil
}
