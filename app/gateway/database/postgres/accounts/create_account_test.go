package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
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
			name: "cadastro com sucesso",
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
		},
		{
			name: "cadastro duplicado",
			input: account.Account{
				ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:       "Joao da Silva",
				Cpf:        "38330499912",
				Balance:    10000,
				Created_at: time.Now(),
			},
			want: account.Account{
				ID:      "",
				Name:    "",
				Cpf:     "",
				Balance: 0,
			},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := accountRepository.Create(ctx, test.input)

			if test.name == "cadastro duplicado" {
				_, err = accountRepository.Create(ctx, test.input)
				got = account.Account{}
			}

			if err == nil {
				test.want.Created_at = got.Created_at
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
