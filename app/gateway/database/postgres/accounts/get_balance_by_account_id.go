package account

import (
	"context"
	"database/sql"
	"errors"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (repository accountRepository) GetBalanceByAccountID(ctx context.Context, accountExternalID types.ExternalID) (types.Money, error) {

	const sqlQuery = `
	SELECT 
		balance
	FROM
		accounts
	WHERE
		external_id = $1
	`

	result := repository.db.QueryRow(
		sqlQuery,
		accountExternalID,
	)
	var balanceValue types.Money
	err := result.Scan(&balanceValue)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, customError.ErrorAccountIDNotFound
		}

		return -1, err
	}

	return balanceValue, nil
}
