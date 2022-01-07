package transfer

import (
	"context"
	"database/sql"
	"stoneBanking/app/domain/entities/transfer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name      string
		runBefore func(db *sql.DB)
		want      int
		wantErr   bool
	}{
		{
			name: "localizados todas as transferencias para conta existente",
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

				input := transfer.Transfer{
					AccountOriginID:      1,
					AccountDestinationID: 1,
					Amount:               100,
					CreatedAt:            time.Now(),
				}
				_, err = transferRepository.Create(ctx, input)

				if err != nil {
					t.Errorf("Não foi possível inicializar os dados iniciais do teste")
				}
			},
			want:    1,
			wantErr: false,
		},
		{
			name:    "teste com o banco vazio, deve retornar lista vazia",
			want:    0,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)

			if test.runBefore != nil {
				test.runBefore(database)
			}

			got, err := transferRepository.GetAll(ctx)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
