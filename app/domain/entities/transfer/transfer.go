package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	ID                     int
	External_ID            types.TransferID
	Account_origin_id      types.AccountOriginID
	Account_destination_id types.AccountDestinyID
	Amount                 types.Money
	Created_at             time.Time
}
