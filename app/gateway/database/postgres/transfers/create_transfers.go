package transfer

import (
	"context"

	"stoneBanking/app/domain/entities/transfer"
)

func (r transferRepository) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	const sqlQueryCreate = `
	INSERT INTO
			transfers (account_origin_id, account_destiny_id, amount)
	VALUES
			($1, $2, $3)
	RETURNING
			id, external_id, created_at
	`

	const sqlQueryUpdateOrigin = `
	UPDATE 
		accounts
	SET 
		balance = (balance - $2)
	WHERE 
		id = $1
	`

	const sqlQueryUpdateDestiny = `
	UPDATE 
		accounts
	SET 
		balance = (balance + $2)
	WHERE 
		id = $1
	`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return transfer.Transfer{}, err
	}

	defer tx.Rollback() //nolint: errorlint
	row := tx.QueryRow(
		sqlQueryCreate,
		transferData.AccountOriginID,
		transferData.AccountDestinationID,
		transferData.Amount)

	err = row.Scan(&transferData.ID, &transferData.ExternalID, &transferData.CreatedAt)
	if err != nil {
		return transfer.Transfer{}, err
	}

	_, err = tx.Exec(sqlQueryUpdateOrigin, transferData.AccountOriginID, transferData.Amount)
	if err != nil {
		return transfer.Transfer{}, err
	}

	_, err = tx.Exec(sqlQueryUpdateDestiny, transferData.AccountDestinationID, transferData.Amount)
	if err != nil {
		return transfer.Transfer{}, err
	}

	err = tx.Commit()
	if err != nil {
		return transfer.Transfer{}, err
	}

	return transferData, nil
}
