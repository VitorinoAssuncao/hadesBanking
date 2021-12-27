package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

func (repository transferRepository) GetByID(ctx context.Context, transferID types.TransferID) (transfer.Transfer, error) {
	var sqlQuery = `
	SELECT 
		id,external_id, account_origin_id, account_destiny_id, amount, created_at
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

	err := result.Scan(&transferData.ID, &transferData.ExternalID, &transferData.AccountOriginID, &transferData.AccountDestinationID, &transferData.Amount, &transferData.CreatedAt)

	if err != nil {
		return transfer.Transfer{}, err
	}

	return transferData, nil
}
