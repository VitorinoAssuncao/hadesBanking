package output

import "stoneBanking/app/domain/entities/transfer"

type TransferOutputVO struct {
	ID                          string  `json:"id" example:"bf89f445-70d3-442f-9821-55699a868704"`
	AccountOriginExternalID     string  `json:"account_origin_id" example:"bf89f445-70d3-442f-9821-55699a868704"`
	AccountDestiationExternalID string  `json:"account_destination_id" example:"bf89f445-70d3-442f-9821-55699a868704"`
	Amount                      float64 `json:"value" example:"123.32"`
	Created_At                  string  `json:"created_at" example:"12/05/2021 00:01:01" `
}

func TransferToTransferOutput(transfer transfer.Transfer) TransferOutputVO {
	transferOutput := TransferOutputVO{
		ID:                          string(transfer.ExternalID),
		AccountOriginExternalID:     string(transfer.AccountOriginExternalID),
		AccountDestiationExternalID: string(transfer.AccountDestinationExternalID),
		Amount:                      transfer.Amount.ToFloat(),
		Created_At:                  transfer.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
	return transferOutput
}
