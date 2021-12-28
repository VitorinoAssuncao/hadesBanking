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
		wantedID  string
		want      int
		wantErr   bool
	}{
		{
			name: "conta localizada, quando usado o id correto",
			input: transfer.Transfer{
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			runBefore: func(db *sql.DB) {
				sqlQuery := `TRUNCATE transfers`
				_, err := db.Exec(sqlQuery)
				if err != nil {
					t.Errorf(err.Error())
				}
			},
			wantedID: "d3280f8c-570a-450d-89f7-3509bc84980d",
			want:     1,
			wantErr:  false,
		},
		{
			name: "conta não localizada, pois id não existe",
			input: transfer.Transfer{
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			wantedID: "d3280f8c-570a-450d-89f7-3509bc849899",
			want:     0,
			wantErr:  false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if test.runBefore != nil {
				test.runBefore(database)
			}
			_, err := transferRepository.Create(ctx, test.input)
			got, err := transferRepository.GetAllByAccountID(ctx, types.AccountID(test.wantedID))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
