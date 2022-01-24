package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
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
		wantErr   bool
	}{
		{
			name: "with the right cpf input, return the account data",
			runBefore: func() (value string) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Balance: 10000,
				}
				created, err := accountRepository.Create(ctx, input)

				if err == nil {
					value = string(created.CPF)
				}

				return value
			},
			want: account.Account{
				ExternalID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:       "Joao da Silva",
				CPF:        "38330499912",
				Balance:    10000,
			},
			wantErr: false,
		},
		{
			name:    "when trying to find a account that not exist, return a error",
			input:   "38330499999",
			want:    account.Account{},
			wantErr: true,
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
				test.want.CreatedAt = got.CreatedAt
				test.want.ExternalID = got.ExternalID
				test.want.ID = got.ID
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
