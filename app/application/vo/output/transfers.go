package output

import "stoneBanking/app/domain/entities/transfer"

type TransferOutputVO struct {
	ID                 string  `json:"transfer_id" example:"1"`
	AccountOriginName  string  `json:"account_origin_name" example:"1"`
	AccountDestinyName string  `json:"account_destiny_name" example:"3"`
	Amount             float64 `json:"value" example:"123.32"`
	Created_At         string  `json:"created_at" example:"12/05/2021 00:01:01" `
}

func TransferToTransferOutput(transfer transfer.Transfer) TransferOutputVO {
	transferOutput := TransferOutputVO{
		ID:                 string(transfer.ExternalID),
		AccountOriginName:  transfer.AccountOriginName,
		AccountDestinyName: transfer.AccountDestinationName,
		Amount:             transfer.Amount.ToFloat(),
		Created_At:         transfer.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
	return transferOutput
}
