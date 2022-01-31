package transfer

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/transfer"
)

func Test_Create(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name      string
		input     transfer.Transfer
		runBefore func(db *sql.DB)
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
			runBefore: func(db *sql.DB) {
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
		t.Run(test.name, func(t *testing.T) {
			if TruncateTable(database) != nil {
				t.Errorf("has not possible clean the databases")
			}

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
