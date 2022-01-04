package transfer

import (
	"context"
	"database/sql"
	"stoneBanking/app/domain/entities/transfer"
	"testing"

	"github.com/stretchr/testify/assert"
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
			name: "conta cadastrada com sucesso, quando dados corretos",
			input: transfer.Transfer{
				AccountOriginID:      "1a05b9b9-6949-40ed-bcfa-aa5c3dd6a88e",
				AccountDestinationID: "7808ae45-ec59-44cd-9458-277564ce7775",
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
				AccountOriginID:      "1a05b9b9-6949-40ed-bcfa-aa5c3dd6a88e",
				AccountDestinationID: "7808ae45-ec59-44cd-9458-277564ce7775",
				Amount:               100,
			},
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
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
