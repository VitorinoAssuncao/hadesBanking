package postgres_transfer

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
			name: "cadastro com sucesso",
			input: transfer.Transfer{
				External_ID:            "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_origin_id:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_destination_id: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:                 100,
				Created_at:             time.Now(),
			},
			want: transfer.Transfer{
				External_ID:            "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_origin_id:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_destination_id: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:                 100,
				Created_at:             time.Now(),
			},
			wantErr: false,
		},
		{
			name: "cadastro duplicado",
			input: transfer.Transfer{
				External_ID:            "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_origin_id:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_destination_id: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:                 100,
				Created_at:             time.Now(),
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
				test.want.Created_at = got.Created_at
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
