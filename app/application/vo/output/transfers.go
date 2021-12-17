package output

import "stoneBanking/app/domain/entities/transfer"

type TransferOutputVO struct {
	ID               string  `json:"transfer_id" example:"1"`
	AccountOriginID  string  `json:"account_origin_id" example:"1"`
	AccountDestinyID string  `json:"account_destiny_id" example:"3"`
	Value            float64 `json:"value" example:"123.32"`
	Created_At       string  `json:"created_at" example:"12/05/2021 00:01:01" `
}

func TransferToTransferOutput(transfer transfer.Transfer) TransferOutputVO {
	transferOutput := TransferOutputVO{
		ID:               string(transfer.ID),
		AccountOriginID:  string(transfer.Account_origin_id),
		AccountDestinyID: string(transfer.Account_destination_id),
		Value:            transfer.Amount.ToFloat(),
		Created_At:       transfer.Created_at.Format("2006-01-02T15:04:05Z07:00"),
	}
	return transferOutput
}
