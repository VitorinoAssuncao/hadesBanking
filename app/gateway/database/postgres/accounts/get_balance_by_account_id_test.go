package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func Test_GetBalanceByAccountID(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		name      string
		want      types.Money
		runBefore func(ac account.Repository) (value types.ExternalID)
		input     types.ExternalID
		wantErr   error
	}{
		{
			name: "with the right input id, return the balance from account",
			runBefore: func(ac account.Repository) (value types.ExternalID) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Balance: 10000,
				}
				created, err := ac.Create(ctx, input)

				if err == nil {
					value = created.ExternalID
				}

				return value
			},
			want:    10000,
			input:   "d3280f8c-570a-450d-89f7-3509bc84980d",
			wantErr: nil,
		}, {
			name:    "when trying to find a account with the wrong id (or the account not exist), return a error and a negative value",
			want:    -1,
			input:   "d3280f8c-570a-450d-89f7-3509bc849899",
			wantErr: customError.ErrorAccountIDNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			database := pgtest.SetDatabase(t, pgtest.GetRandomDBName())
			accountRepository := NewAccountRepository(database)

			if test.runBefore != nil {
				test.input = test.runBefore(accountRepository)
			}
			got, err := accountRepository.GetBalanceByAccountID(ctx, types.ExternalID(test.input))

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
