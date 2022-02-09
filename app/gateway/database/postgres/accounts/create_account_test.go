package account

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func Test_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	testCases := []struct {
		name      string
		input     account.Account
		runBefore func(db *pgxpool.Pool)
		want      account.Account
		wantErr   bool
	}{
		{
			name: "with right data, account is created successfully",
			input: account.Account{
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: time.Now(),
			},
			want: account.Account{
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			wantErr: false,
		},
		{
			name: "when trying to create a account duplicated, return a error",
			input: account.Account{
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: time.Now(),
			},
			runBefore: func(db *pgxpool.Pool) {
				sqlQuery :=
					`
				INSERT INTO
					accounts (name, cpf, secret, balance)
				VALUES
					('Joao da Silva', '38330499912', 'password', 100)
				`
				_, err := db.Exec(ctx, sqlQuery)
				if err != nil {
					t.Errorf(err.Error())
				}
			},
			want:    account.Account{},
			wantErr: true,
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

			if test.runBefore != nil {
				test.runBefore(database)
			}
			got, err := accountRepository.Create(ctx, test.input)

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
