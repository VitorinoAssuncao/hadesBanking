package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                   int
	ExternalID           types.TransferID
	AccountOriginID      types.AccountTransferID
	AccountDestinationID types.AccountTransferID
	Amount               types.Money
	CreatedAt            time.Time
}
