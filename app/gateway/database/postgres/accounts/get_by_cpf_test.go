package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetByCPF(t *testing.T) {
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
			name: "localizada a conta utilizando-se do cpf, e retorna os dados da mesma",
			input: account.Account{
				ID:        "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: time.Now(),
			},
			want: account.Account{
				ID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			wantErr: false,
		},
		{
			name: "retorna erro ao tentar localizar conta com cpf inexistente",
			input: account.Account{
				ID:        "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: time.Now(),
			},
			want: account.Account{
				ID:      "",
				Name:    "",
				CPF:     "38330499999",
				Balance: 0,
			},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := accountRepository.Create(ctx, test.input)
			got, err := accountRepository.GetByCPF(ctx, test.want.CPF)
			if err == nil {
				test.want.CreatedAt = got.CreatedAt
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want.Name, got.Name)
		})
	}
}
