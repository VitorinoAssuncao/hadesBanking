package transfer

import (
	"context"
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
		input     transfer.Transfer
		runBefore func(value string) string
		wantedID  string
		want      transfer.Transfer
		wantErr   bool
	}{
		{
			name: "conta localizada com sucesso, retorna dados da conta",
			input: transfer.Transfer{
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			runBefore: func(value string) string {
				return value
			},
			wantedID: "d3280f8c-570a-450d-89f7-3509bc84980d",
			want: transfer.Transfer{
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			wantErr: false,
		},
		{
			name: "busca por conta inexistente, deve retornar erro e dados",
			input: transfer.Transfer{
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			wantedID: "d3280f8c-570a-450d-89f7-3509bc849899",
			want:     transfer.Transfer{},
			wantErr:  true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			created, err := transferRepository.Create(ctx, test.input)

			if err != nil {
				t.Errorf(err.Error())
			}

			if test.runBefore != nil {
				test.wantedID = test.runBefore(string(created.ExternalID))
			}

			got, err := transferRepository.GetByID(ctx, types.TransferID(test.wantedID))

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
