package account

import (
	"context"
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
		return 0, err
	}

	return balanceValue, nil
}
