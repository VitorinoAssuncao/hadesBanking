package account

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Account struct {
	ID         int
	ExternalID types.ExternalID
	Name       string
	CPF        types.Document
	Secret     types.Password
	Balance    types.Money
	CreatedAt  time.Time
}
