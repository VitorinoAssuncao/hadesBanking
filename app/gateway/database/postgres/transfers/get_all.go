package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
)

func (repository transferRepository) GetAll(ctx context.Context) ([]transfer.Transfer, error) {
	var transfers = make([]transfer.Transfer, 0)
	const sqlQuery = `
	SELECT 
		id,external_id, account_origin_id, account_destiny_id, amount, created_at
	FROM
		transfers
	`
	result, err := repository.db.Query(sqlQuery)
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
