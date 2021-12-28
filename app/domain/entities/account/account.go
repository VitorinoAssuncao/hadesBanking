package account

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Account struct {
	ID        types.AccountID
	Name      string
	CPF       string
	Secret    string
	Balance   types.Money
	CreatedAt time.Time
}
