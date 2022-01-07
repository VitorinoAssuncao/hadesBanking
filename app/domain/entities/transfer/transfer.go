package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                           types.InternalID
	ExternalID                   types.ExternalID
	AccountOriginID              types.InternalID
	AccountOriginExternalID      types.ExternalID
	AccountDestinationID         types.InternalID
	AccountDestinationExternalID types.ExternalID
	Amount                       types.Money
	CreatedAt                    time.Time
}
