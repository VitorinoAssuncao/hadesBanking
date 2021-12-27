package postgres_transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
)

func (repository transferRepository) GetAll(ctx context.Context) ([]transfer.Transfer, error) {
	var tempTransfer transfer.Transfer
	var transfers = make([]transfer.Transfer, 0)
	var sqlQuery = `
	SELECT 
		id,external_id, account_origin_id, account_destiny_id, amount, created_at
	FROM
		transfers
	`
	result, err := repository.db.Query(sqlQuery)
	if err != nil {
		return transfers, err
	}

	for result.Next() {
		err = result.Scan(&tempTransfer.ID, &tempTransfer.External_ID, &tempTransfer.Account_origin_id, &tempTransfer.Account_destination_id, &tempTransfer.Amount, &tempTransfer.Created_at)
		if err != nil {
			return transfers, err
		}
		transfers = append(transfers, tempTransfer)
	}

	return transfers, nil
}
