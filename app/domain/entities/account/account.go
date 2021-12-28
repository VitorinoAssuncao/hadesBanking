package account

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Account struct {
	ID         types.AccountID
	Name       string
	Cpf        string
	Secret     string
	Balance    types.Money
	Created_at time.Time
}
