package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                   int
	ExternalID           types.TransferExternalID
	AccountOriginID      types.TransferAccountID
	AccountDestinationID types.TransferAccountID
	Amount               types.Money
	CreatedAt            time.Time
}
