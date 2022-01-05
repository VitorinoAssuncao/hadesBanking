package account

import (
	"context"
	"stoneBanking/app/domain/types"
)

func (repository accountRepository) UpdateBalance(ctx context.Context, value int, external_id types.AccountExternalID) (bool, error) {
	const sqlQuery = `
	UPDATE
			accounts
	SET
			balance = $1
	WHERE
			external_id = $2
	`
	_, err := repository.db.Exec(sqlQuery, value, external_id)

	if err != nil {
		return false, err
	}
	return true, nil
}