package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

func (r transferRepository) GetAllByAccountID(ctx context.Context, acccountID types.AccountExternalID) ([]transfer.Transfer, error) {
	var transfers = make([]transfer.Transfer, 0)
	const sqlQuery = `
	SELECT 
		id,external_id, account_origin_id, account_destiny_id, amount, created_at
	FROM
		transfers
	WHERE
		account_origin_id = $1 or account_destiny_id = $1
	`
	result, err := r.db.Query(sqlQuery, acccountID)
	if err != nil {
		return transfers, err
	}

	var tempTransfer transfer.Transfer

	for result.Next() {
		err = result.Scan(&tempTransfer.ID, &tempTransfer.ExternalID, &tempTransfer.AccountOriginID, &tempTransfer.AccountDestinationID, &tempTransfer.Amount, &tempTransfer.CreatedAt)
		if err != nil {
			return transfers, err
		}
		transfers = append(transfers, tempTransfer)
	}

	return transfers, nil
}
