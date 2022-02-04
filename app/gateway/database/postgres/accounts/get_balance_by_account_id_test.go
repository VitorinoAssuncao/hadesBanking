package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func Test_GetBalanceByAccountID(t *testing.T) {
	ctx := context.Background()
	database := &testConn
	accountRepository := NewAccountRepository(*database)
	testCases := []struct {
		name      string
		want      types.Money
		runBefore func() (value types.ExternalID)
		input     types.ExternalID
		wantErr   error
	}{
		{
			name: "with the right input id, return the balance from account",
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
			if TruncateTable(ctx, *database) != nil {
				t.Errorf("has not possible clean the databases")
			}

			if test.runBefore != nil {
				test.input = test.runBefore()
			}
			got, err := accountRepository.GetBalanceByAccountID(ctx, types.ExternalID(test.input))

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
