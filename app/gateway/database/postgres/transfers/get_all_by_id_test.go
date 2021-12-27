package transfer

import (
	"context"
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
		name    string
		input   transfer.Transfer
		want    int
		wantErr bool
	}{
		{
			name: "conta localizada, quando usado o id correto",
			input: transfer.Transfer{
				External_ID:            "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_origin_id:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_destination_id: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:                 100,
				Created_at:             time.Now(),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "conta não localizada, pois id não existe",
			input: transfer.Transfer{
				External_ID:            "d3280f8c-570a-450d-89f7-3509bc849899",
				Account_origin_id:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Account_destination_id: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Amount:                 100,
				Created_at:             time.Now(),
			},
			want:    0,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := transferRepository.Create(ctx, test.input)
			got, err := transferRepository.GetAllByAccountID(ctx, types.AccountID(test.input.External_ID))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
