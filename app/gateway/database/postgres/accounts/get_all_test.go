package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name    string
		input   account.Account
		want    int
		wantErr bool
	}{
		{
			name: "localizar todas as contas",
			input: account.Account{
				ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:       "Joao da Silva",
				Cpf:        "38330499912",
				Balance:    10000,
				Created_at: time.Now(),
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := accountRepository.Create(ctx, test.input)
			got, err := accountRepository.GetAll(ctx)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
