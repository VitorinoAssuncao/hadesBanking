package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetByCPF(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name      string
		input     string
		runBefore func() (value string)
		want      account.Account
		wantErr   error
	}{
		{
			name: "with the right cpf input, return the account data",
			runBefore: func() (value string) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Secret:  "12345",
					Balance: 10000,
				}
				created, err := accountRepository.Create(ctx, input)

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
		t.Run(test.name, func(t *testing.T) {
			if TruncateTable(database) != nil {
				t.Errorf("has not possible clean the databases")
			}

			if test.runBefore != nil {
				test.input = test.runBefore()
			}

			got, err := accountRepository.GetByCPF(ctx, test.input)

			if err == nil {
				test.want.ExternalID = got.ExternalID
			}

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
