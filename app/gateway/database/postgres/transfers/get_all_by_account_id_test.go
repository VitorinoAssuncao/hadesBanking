package transfer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func Test_GetAllByID(t *testing.T) {
	ctx := context.Background()
	database := testPool
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name      string
		runBefore func(db *pgxpool.Pool) (value types.InternalID)
		input     int
		want      []transfer.Transfer
		wantErr   error
	}{
		{
			name: "with a valid id in the input, find the transfers and return without errors",

			runBefore: func(db *pgxpool.Pool) (value types.InternalID) {
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

				input := transfer.Transfer{
					AccountOriginID:      1,
					AccountDestinationID: 1,
					Amount:               100,
					CreatedAt:            time.Now(),
				}
				created, err := transferRepository.Create(ctx, input)
				if err != nil {
					t.Errorf("has not possible initialize the test data")
				}
				return created.ID
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
		t.Run(test.name, func(t *testing.T) {

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
