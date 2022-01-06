package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                           types.InternalID
	ExternalID                   types.TransferExternalID
	AccountOriginID              types.InternalID
	AccountOriginExternalID      types.AccountExternalID
	AccountOriginName            string
	AccountDestinationID         types.InternalID
	AccountDestinationExternalID types.AccountExternalID
	AccountDestinationName       string
	Amount                       types.Money
	CreatedAt                    time.Time
}
