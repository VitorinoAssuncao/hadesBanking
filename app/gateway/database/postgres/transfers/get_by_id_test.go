package transfer

import (
	"context"
	"database/sql"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name      string
		runBefore func(db *sql.DB) string
		input     string
		want      transfer.Transfer
		wantErr   bool
	}{
		{
			name: "with the right id, locate the account and return withouth error",
			runBefore: func(db *sql.DB) (value string) {
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
					Amount:               100,
					AccountOriginID:      1,
					AccountDestinationID: 1,
					CreatedAt:            time.Now(),
				}
				created, err := transferRepository.Create(ctx, input)

				if err != nil {
					t.Errorf("has not possible initialize the test data")
				}

				return string(created.ExternalID)
			},
			want: transfer.Transfer{
				Amount:               100,
				AccountOriginID:      1,
				AccountDestinationID: 1,
				CreatedAt:            time.Now(),
			},
			wantErr: false,
		},
		{
			name:    "try to find a account with a id that not exist, returning a error",
			input:   "d3280f8c-570a-450d-89f7-3509bc849899",
			want:    transfer.Transfer{},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if TruncateTable(database) != nil {
				t.Errorf("has not possible clean the databases")
			}

			if test.runBefore != nil {
				test.input = test.runBefore(database)
			}

			got, err := transferRepository.GetByID(ctx, types.ExternalID(test.input))

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
