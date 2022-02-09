package transfer

import (
	"context"
	"log"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func Test_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	testCases := []struct {
		name      string
		input     transfer.Transfer
		runBefore func(db *pgxpool.Pool)
		want      transfer.Transfer
		wantErr   bool
	}{
		{
			name: "with the right input data, create the transfer successfully",
			input: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 1,
				Amount:               100,
			},
			runBefore: func(db *pgxpool.Pool) {
				acc := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Secret:  "password",
					Balance: 100,
				}
				_, err := pgtest.CreateAccount(db, acc)
				if err != nil {
					t.Errorf("was not possible to create the test account %s", err.Error())
				}
			},
			want: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 1,
				Amount:               100,
			},
			wantErr: false,
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

			transferRepository := NewTransferRepository(database)

			if test.runBefore != nil {
				test.runBefore(database)
			}
			got, err := transferRepository.Create(ctx, test.input)

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
