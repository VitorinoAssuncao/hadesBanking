package postgres_transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
)

func (repository transferRepository) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	var sqlQuery = `
	INSERT INTO
			transfers (external_id, account_origin_id, account_destiny_id, amount, created_at)
	VALUES
			($1, $2, $3, $4, $5)
	`
	_, err := repository.db.Exec(
		sqlQuery,
		transferData.External_ID,
		transferData.Account_origin_id,
		transferData.Account_destination_id,
		transferData.Amount,
		transferData.Created_at)

	if err != nil {
		return transfer.Transfer{}, err
	}
	return transferData, nil
}
