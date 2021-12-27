package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/transfer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
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
			name: "conta cadastrada com sucesso, quando dados corretos",
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
			name: "cadastro gera erro, quando tentativa de criar conta com cpf existente",
			input: transfer.Transfer{
				ExternalID:           "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountOriginID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				AccountDestinationID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:               100,
				CreatedAt:            time.Now(),
			},
			want:    transfer.Transfer{},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := transferRepository.Create(ctx, test.input)

			if test.name == "cadastro duplicado" {
				_, err = transferRepository.Create(ctx, test.input)
				got = transfer.Transfer{}
			}

			if err == nil {
				test.want.CreatedAt = got.CreatedAt
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
