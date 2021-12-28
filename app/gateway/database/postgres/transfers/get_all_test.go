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
		input     transfer.Transfer
		runBefore func(db *sql.DB)
		want      int
		wantErr   bool
	}{
		{
			name: "localizados todas as transferencias para conta existente",
			input: transfer.Transfer{
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			runBefore: func(db *sql.DB) {
				truncateQuery := `TRUNCATE transfers`
				_, err := db.Exec(truncateQuery)

				if err != nil {
					t.Errorf(err.Error())
				}

				sqlQuery := `
				INSERT INTO
					transfers (external_id, account_origin_id, account_destiny_id, amount, created_at)
				VALUES
					('d3280f8c-570a-450d-89f7-3509bc84980d', 'd3280f8c-570a-450d-89f7-3509bc84980d', 'd3280f8c-570a-450d-89f7-3509bc84980d', 100, $1)
				`
				_, err = db.Exec(sqlQuery, time.Now())

				if err != nil {
					t.Errorf(err.Error())
				}
			},
			want:    1,
			wantErr: false,
		},
		{
			name:  "teste com o banco vazio, deve retornar lista vazia",
			input: transfer.Transfer{},
			runBefore: func(db *sql.DB) {
				sqlQuery := `TRUNCATE transfers`
				_, err := db.Exec(sqlQuery)
				if err != nil {
					t.Errorf(err.Error())
				}
			},
			want:    0,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if test.runBefore != nil {
				test.runBefore(database)
			}

			got, err := transferRepository.GetAll(ctx)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
