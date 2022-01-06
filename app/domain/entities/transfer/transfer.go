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
	AccountDestinationID         types.InternalID
	AccountDestinationExternalID types.AccountExternalID
	Amount                       types.Money
	CreatedAt                    time.Time
}
