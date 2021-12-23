package postgres_transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
)

func (repository transferRepository) Create(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
	var sqlQuery = `
	INSERT INTO
			transfers (id,external_id, account_origin_id, account_destiny_id, ammount, created_at)
	VALUES
			($1, $2, $3, $4, $5, $6)
	`
	_, err := repository.db.Exec(
		sqlQuery,
		transferData.ID,
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

func (repository transferRepository) GetAll(ctx context.Context) ([]transfer.Transfer, error) {
	transfers := []transfer.Transfer{}
	return transfers, nil
}
