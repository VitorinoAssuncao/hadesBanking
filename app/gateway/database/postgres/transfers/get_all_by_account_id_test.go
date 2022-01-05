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
			name: "conta localizada, quando usado o id correto",

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
					t.Errorf("Não foi possível inicializar o teste")
				}
				return created.ID
			},
			want:    1,
			wantErr: false,
		},
		{
			name:    "conta não localizada, pois id não existe",
			input:   91,
			want:    0,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
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
