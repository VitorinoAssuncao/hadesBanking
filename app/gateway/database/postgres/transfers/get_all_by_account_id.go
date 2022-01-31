package transfer

import (
	"context"

	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func (r transferRepository) GetAllByAccountID(ctx context.Context, acccountID types.InternalID) ([]transfer.Transfer, error) {
	var transfers = make([]transfer.Transfer, 0)
	const sqlQuery = `
	SELECT
		t.id, t.external_id, t.account_origin_id, t.account_destiny_id, t.amount, t.created_at,o.external_id ,d.external_id
	FROM
		transfers t
	INNER JOIN
		accounts o ON (o.id = t.account_origin_id )
	INNER JOIN
		accounts d on (d.id = t.account_destiny_id)
	WHERE
		(t.account_origin_id = $1 or t.account_destiny_id = $1)`

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
			&tempTransfer.AccountDestinationExternalID,
		)
		if err != nil {
			return transfers, err
		}

		transfers = append(transfers, tempTransfer)
	}

	if len(transfers) <= 0 {
		return []transfer.Transfer{}, customError.ErrorTransferAccountNotFound
	}

	return transfers, nil
}
