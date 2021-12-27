package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                   int
	ExternalID           types.TransferID
	AccountOriginID      types.AccountOriginID
	AccountDestinationID types.AccountDestinyID
	Amount               types.Money
	CreatedAt            time.Time
}
