package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	transferRepository := NewTransferRepository(database)
	testCases := []struct {
		name    string
		input   transfer.Transfer
		want    transfer.Transfer
		wantErr bool
	}{
		{
			name: "conta localizada com sucesso, retorna dados da conta",
			input: transfer.Transfer{
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			want: transfer.Transfer{
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
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
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			want:    transfer.Transfer{},
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := transferRepository.Create(ctx, test.input)

			if err != nil {
				t.Errorf(err.Error())
			}

			got, err := transferRepository.GetByID(ctx, test.want.ExternalID)

			if err == nil {
				test.want.CreatedAt = got.CreatedAt
				test.want.ID = got.ID
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
