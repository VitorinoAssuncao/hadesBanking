package transfer

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func Test_GetAllByID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	testCases := []struct {
		name      string
		runBefore func(db *pgxpool.Pool) (value types.ExternalID)
		input     int
		want      []transfer.Transfer
		wantErr   error
	}{
		{
			name: "with a valid id in the input, find the transfers and return without errors",

			runBefore: func(db *pgxpool.Pool) (value types.ExternalID) {
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

				sqlQuery2 := `
				INSERT INTO
					transfers (account_origin_id, account_destiny_id, amount)
				VALUES
					(1, 1, 100)
				RETURNING
					id
				`
				created, err := db.Exec(ctx, sqlQuery2)
				if err != nil {
					t.Errorf(err.Error())
				}
				return types.ExternalID(created)
			},
			want: []transfer.Transfer{
				{
					AccountOriginID:      1,
					AccountDestinationID: 1,
					Amount:               100,
					CreatedAt:            time.Now(),
				},
			},
			wantErr: nil,
		},
		{
			name:    "with a id that not exist, return a error that the account not exist",
			input:   91,
			want:    []transfer.Transfer{},
			wantErr: customError.ErrorTransferAccountNotFound,
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
				fmt.Println("valor:", test.runBefore(database))
				test.input = 1
			}
			got, err := transferRepository.GetAllByAccountID(ctx, types.InternalID(test.input))
			for index, result := range got {
				test.want[index].ID = result.ID
				test.want[index].ExternalID = result.ExternalID
				test.want[index].AccountOriginExternalID = result.AccountOriginExternalID
				test.want[index].AccountDestinationExternalID = result.AccountDestinationExternalID
				test.want[index].CreatedAt = result.CreatedAt
			}

			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}
