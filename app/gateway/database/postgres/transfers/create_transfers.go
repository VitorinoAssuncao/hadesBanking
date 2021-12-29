package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
)

func (r transferRepository) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	const sqlQuery = `
	INSERT INTO
			transfers (account_origin_id, account_destiny_id, amount)
	VALUES
			($1, $2, $3)
	RETURNING
			id, external_id, created_at
	`
	row := r.db.QueryRow(
		sqlQuery,
		transferData.AccountOriginID,
		transferData.AccountDestinationID,
		transferData.Amount)

	err := row.Scan(&transferData.ID, &transferData.ExternalID, &transferData.CreatedAt)

	if err != nil {
		return transfer.Transfer{}, err
	}
	return transferData, nil
}
