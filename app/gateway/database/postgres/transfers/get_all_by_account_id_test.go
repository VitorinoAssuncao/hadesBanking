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
		input     transfer.Transfer
		runBefore func(db *sql.DB)
		wantedID  int
		want      int
		wantErr   bool
	}{
		{
			name: "conta localizada, quando usado o id correto",
			input: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 1,
				Amount:               100,
				CreatedAt:            time.Now(),
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
			wantedID: 1,
			want:     1,
			wantErr:  false,
		},
		{
			name:     "conta não localizada, pois id não existe",
			input:    transfer.Transfer{},
			wantedID: 99,
			want:     0,
			wantErr:  false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			if test.runBefore != nil {
				test.runBefore(database)
			}
			_, err := transferRepository.Create(ctx, test.input)
			got, err := transferRepository.GetAllByAccountID(ctx, types.AccountTransferID(test.wantedID))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
