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

func Test_GetAllByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name      string
		runBefore func(db *sql.DB) (value string)
		input     string
		want      int
		wantErr   bool
	}{
		{
			name: "conta localizada, quando usado o id correto",

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
					AccountOriginID:      "1a05b9b9-6949-40ed-bcfa-aa5c3dd6a88e",
					AccountDestinationID: "7808ae45-ec59-44cd-9458-277564ce7775",
					Amount:               100,
					CreatedAt:            time.Now(),
				}
				created, err := transferRepository.Create(ctx, input)
				return string(created.ExternalID)
			},
			want:    1,
			wantErr: false,
		},
		{
			name:    "conta não localizada, pois id não existe",
			input:   "9191919191",
			want:    0,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			if test.runBefore != nil {
				test.input = test.runBefore(database)
			}
			got, err := transferRepository.GetAllByAccountID(ctx, types.AccountExternalID(test.input))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
