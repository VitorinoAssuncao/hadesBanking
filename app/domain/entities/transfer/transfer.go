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
	AccountOriginName            string
	AccountDestinationID         types.InternalID
	AccountDestinationExternalID types.ExternalID
	AccountDestinationName       string
	Amount                       types.Money
	CreatedAt                    time.Time
}
