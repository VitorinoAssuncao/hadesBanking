package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
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
		wantErr   error
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
			wantErr: nil,
		}, {
			name:    "when trying to find a account with the wrong id (or the account not exist), return a error and a void object",
			want:    account.Account{},
			input:   "d3280f8c-570a-450d-89f7-3509bc849899",
			wantErr: customError.ErrorAccountIDNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if TruncateTable(ctx, database) != nil {
				t.Errorf("has not possible clean the databases")
			}

			if test.runBefore != nil {
				test.input = test.runBefore()
			}
			got, err := accountRepository.GetByID(ctx, types.ExternalID(test.input))
			if err == nil {
				test.want.CreatedAt = got.CreatedAt
				test.want.ID = got.ID
				test.want.ExternalID = got.ExternalID
			}
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
