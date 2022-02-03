package account

import (
	"context"
	"errors"

	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"

	"github.com/jackc/pgx/v4"
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
		ctx,
		sqlQuery,
		accountExternalID.ToUUID(),
	)

	var balanceValue types.Money
	err := result.Scan(&balanceValue)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return -1, customError.ErrorAccountIDNotFound
		}

		return -1, err
	}

	return balanceValue, nil
}
