package transfer

import (
	"context"
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
		name    string
		input   transfer.Transfer
		want    int
		wantErr bool
	}{
		{
			name: "localizados todas as transferencias para conta existente",
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
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := transferRepository.Create(ctx, test.input)
			got, err := transferRepository.GetAll(ctx)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
