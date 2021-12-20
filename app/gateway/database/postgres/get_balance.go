package postgres

import (
	"context"
	"stoneBanking/app/domain/types"
)

func (repository accountRepository) GetBalance(ctx context.Context, accountID types.AccountID) (types.Money, error) {

	var sqlQuery = `
	SELECT 
		 balance
	FROM
		accounts
	WHERE
			id = $1
	`
	result := repository.db.QueryRow(
		ctx,
		sqlQuery,
		accountID,
	)
	var value types.Money
	err := result.Scan(&value)

	if err != nil {
		return 0, errorAccountIDNotFound
	}

	return value, nil
}
