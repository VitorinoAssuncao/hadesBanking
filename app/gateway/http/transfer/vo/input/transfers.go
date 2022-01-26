package input

import (
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
)

type CreateTransferVO struct {
	AccountOriginID      string `json:"account_origin_id" example:"2"`
	AccountDestinationID string `json:"account_destination_id" example:"3"`
	Amount               int    `json:"amount" example:"1000"`
}

func (inputTransfer CreateTransferVO) ToEntitie() transfer.Transfer {
	transfer := transfer.Transfer{
		AccountOriginExternalID:      types.ExternalID(inputTransfer.AccountOriginID),
		AccountDestinationExternalID: types.ExternalID(inputTransfer.AccountDestinationID),
		Amount:                       types.Money(inputTransfer.Amount),
	}
	return transfer
}
