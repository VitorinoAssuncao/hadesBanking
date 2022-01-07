package input

import (
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

type CreateTransferVO struct {
	AccountOriginID  string `json:"account_origin_id" example:"2"`
	AccountDestinyID string `json:"account_destiny_id" example:"3"`
	Amount           int    `json:"amount" example:"1000"`
}

func GenerateTransfer(input CreateTransferVO) transfer.Transfer {
	transfer := transfer.Transfer{
		AccountOriginExternalID:      types.ExternalID(input.AccountOriginID),
		AccountDestinationExternalID: types.ExternalID(input.AccountDestinyID),
		Amount:                       types.Money(input.Amount),
	}
	return transfer
}
