package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

func (r transferRepository) GetAllByAccountID(ctx context.Context, acccountID types.InternalID) ([]transfer.Transfer, error) {
	var transfers = make([]transfer.Transfer, 0)
	const sqlQuery = `
	SELECT
		t.id, t.external_id, t.account_origin_id, t.account_destiny_id, t.amount, t.created_at,o.external_id, o.name ,d.external_id, d.name
	FROM
		transfers t
	INNER JOIN
		accounts o ON (o.id = t.account_origin_id )
	INNER JOIN
		accounts d on (d.id = t.account_destiny_id)
	WHERE
		(account_origin_id = $1 or account_destiny_id = $1)`

	result, err := r.db.Query(sqlQuery, acccountID)
	if err != nil {
		return transfers, err
	}

	var tempTransfer transfer.Transfer

	for result.Next() {
		err = result.Scan(
			&tempTransfer.ID,
			&tempTransfer.ExternalID,
			&tempTransfer.AccountOriginID,
			&tempTransfer.AccountDestinationID,
			&tempTransfer.Amount,
			&tempTransfer.CreatedAt,
			&tempTransfer.AccountOriginExternalID,
			&tempTransfer.AccountOriginName,
			&tempTransfer.AccountDestinationExternalID,
			&tempTransfer.AccountDestinationName)
		if err != nil {
			return transfers, err
		}

		transfers = append(transfers, tempTransfer)
	}

	return transfers, nil
}
