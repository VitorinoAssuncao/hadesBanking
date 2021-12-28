package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name    string
		input   account.Account
		want    account.Account
		wantErr bool
	}{
		{
			name: "localizado a conta usando o ID",
			input: account.Account{
				ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:       "Joao da Silva",
				Cpf:        "38330499912",
				Balance:    10000,
				Created_at: time.Now(),
			},
			want: account.Account{
				ID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:    "Joao da Silva",
				Cpf:     "38330499912",
				Balance: 10000,
			},
			wantErr: false,
		}, {
			name: "tentar localizar conta que não existe",
			input: account.Account{
				ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:       "Joao da Silva",
				Cpf:        "38330499912",
				Balance:    10000,
				Created_at: time.Now(),
			},
			want: account.Account{
				ID:      "d3280f8c-570a-450d-89f7-3509bc849899",
				Name:    "",
				Cpf:     "",
				Balance: 0,
			},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := accountRepository.Create(ctx, test.input)
			got, err := accountRepository.GetByID(ctx, test.want.ID)
			if err == nil {
				test.want.Created_at = got.Created_at
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want.Name, got.Name)
		})
	}
}
