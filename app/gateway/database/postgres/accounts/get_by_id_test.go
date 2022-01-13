package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name      string
		want      account.Account
		runBefore func() (value types.ExternalID)
		input     types.ExternalID
		wantErr   bool
	}{
		{
			name: "with the right input id, return the data from account",
			runBefore: func() (value types.ExternalID) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Balance: 10000,
				}
				created, err := accountRepository.Create(ctx, input)

				if err == nil {
					value = created.ExternalID
				}

				return value
			},
			want: account.Account{
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			input:   "d3280f8c-570a-450d-89f7-3509bc84980d",
			wantErr: false,
		}, {
			name:    "when trying to find a account with the wrong id (or the account not exist), return a error and a void object",
			want:    account.Account{},
			input:   "d3280f8c-570a-450d-89f7-3509bc849899",
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			if test.runBefore != nil {
				test.input = test.runBefore()
			}
			got, err := accountRepository.GetByID(ctx, types.ExternalID(test.input))
			if err == nil {
				test.want.CreatedAt = got.CreatedAt
				test.want.ID = got.ID
				test.want.ExternalID = got.ExternalID
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
