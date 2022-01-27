package transfer

import (
	"context"
	"database/sql"
	"fmt"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name      string
		runBefore func(db *sql.DB) (value types.InternalID)
		input     int
		want      int
		wantErr   bool
	}{
		{
			name: "with a valid id in the input, find the transfers and return without errors",

			runBefore: func(db *sql.DB) (value types.InternalID) {
				sqlQuery :=
					`
				INSERT INTO
					accounts (name, cpf, secret, balance)
				VALUES
					('Joao da Silva', '38330499912', 'password', 100)
				`
				_, err := db.Exec(sqlQuery)
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
			want:    1,
			wantErr: false,
		},
		{
			name:    "with a id that not exist, return a error that the account not exist",
			input:   91,
			want:    0,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if TruncateTable(database) != nil {
				t.Errorf("has not possible clean the databases")
			}

			if test.runBefore != nil {
				fmt.Println("valor:", test.runBefore(database))
				test.input = 1
			}
			got, err := transferRepository.GetAllByAccountID(ctx, types.InternalID(test.input))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
