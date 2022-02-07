package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func Test_GetCredentialByCPF(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	testCases := []struct {
		name      string
		input     string
		runBefore func(ac account.Repository) (value string)
		want      account.Account
		wantErr   error
	}{
		{
			name: "with the right cpf input, return the account data",
			runBefore: func(ac account.Repository) (value string) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Secret:  "12345",
					Balance: 10000,
				}
				created, err := ac.Create(ctx, input)

				if err == nil {
					value = string(created.CPF)
				}

				return value
			},
			want: account.Account{
				CPF:    "38330499912",
				Secret: "12345",
			},
			wantErr: nil,
		},
		{
			name:    "when trying to find a account that not exist, return a error",
			input:   "38330499999",
			want:    account.Account{},
			wantErr: customError.ErrorAccountCPFNotFound,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			database, teardown := pgtest.SetDatabase(t, pgtest.GetRandomDBName())
			defer teardown()
			accountRepository := NewAccountRepository(database)

			if test.runBefore != nil {
				test.input = test.runBefore(accountRepository)
			}

			got, err := accountRepository.GetCredentialByCPF(ctx, test.input)

			if err == nil {
				test.want.ExternalID = got.ExternalID
			}

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
