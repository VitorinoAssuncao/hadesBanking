package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                   types.InternalID
	ExternalID           types.TransferExternalID
	AccountOriginID      types.InternalID
	AccountDestinationID types.InternalID
	Amount               types.Money
	CreatedAt            time.Time
}
