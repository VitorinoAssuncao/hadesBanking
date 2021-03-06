package account

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func Test_GetAll(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	testCases := []struct {
		name    string
		input   account.Account
		want    []account.Account
		wantErr error
	}{
		{
			name: "find and return all the accounts, when at last one exist",
			input: account.Account{
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			want: []account.Account{
				{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Balance: 10000,
				},
			},
			wantErr: nil,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			database, err := pgtest.SetDatabase(pgtest.GetRandomDBName())
			if err != nil {
				log.Fatalf(err.Error())
			}
			accountRepository := NewAccountRepository(database)

			_, err = accountRepository.Create(ctx, test.input)
			if err != nil {
				t.Errorf("error when creating account")
			}

			got, err := accountRepository.GetAll(ctx)
			if err == nil {
				for index, result := range got {
					test.want[index].ID = result.ID
					test.want[index].CreatedAt = result.CreatedAt
					test.want[index].ExternalID = result.ExternalID
				}
			}

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
