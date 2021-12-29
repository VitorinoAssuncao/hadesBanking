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
			name: "conta localizada com sucesso, retorna dados da conta",
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
					AccountOriginID:      1,
					AccountDestinationID: 1,
					Amount:               100,
					CreatedAt:            time.Now(),
				}
				created, err := transferRepository.Create(ctx, input)
				return string(created.ExternalID)
			},
			want: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 1,
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			wantErr: false,
		},
		{
			name:    "busca por conta inexistente, deve retornar erro e dados",
			input:   "d3280f8c-570a-450d-89f7-3509bc849899",
			want:    transfer.Transfer{},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			if test.runBefore != nil {
				test.input = test.runBefore(database)
			}

			got, err := transferRepository.GetByID(ctx, types.TransferID(test.input))

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
