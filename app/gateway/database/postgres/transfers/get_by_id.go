package postgres_transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

func (repository transferRepository) GetByID(ctx context.Context, transferID types.TransferID) (transfer.Transfer, error) {
	var sqlQuery = `
	SELECT 
		id,external_id, account_origin_id, account_destiny_id, ammount, created_at
	FROM
		transfers
	WHERE
		external_id = $1
	`

	result := repository.db.QueryRow(
		sqlQuery,
		transferID,
	)
	transferData := transfer.Transfer{}

	err := result.Scan(&transferData.ID, &transferData.External_ID, &transferData.Account_origin_id, &transferData.Account_destination_id, &transferData.Amount, &transferData.Created_at)

	if err != nil {
		return transfer.Transfer{}, err
	}

	return transferData, nil
}
