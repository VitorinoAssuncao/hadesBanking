package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                   int
	ExternalID           types.TransferExternalID
	AccountOriginID      types.AccountExternalID
	AccountDestinationID types.AccountExternalID
	Amount               types.Money
	CreatedAt            time.Time
}
