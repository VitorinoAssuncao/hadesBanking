package account

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Account struct {
	ID         int
	ExternalID types.ExternalID
	Name       string
	CPF        string
	Secret     string
	Balance    types.Money
	CreatedAt  time.Time
}
