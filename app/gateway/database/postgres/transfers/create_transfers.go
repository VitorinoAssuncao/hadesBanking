package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
)

func (r transferRepository) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	const sqlQuery = `
	INSERT INTO
			transfers (external_id, account_origin_id, account_destiny_id, amount, created_at)
	VALUES
			($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(
		sqlQuery,
		transferData.ExternalID,
		transferData.AccountOriginID,
		transferData.AccountDestinationID,
		transferData.Amount,
		transferData.CreatedAt)

	if err != nil {
		return transfer.Transfer{}, err
	}
	return transferData, nil
}
