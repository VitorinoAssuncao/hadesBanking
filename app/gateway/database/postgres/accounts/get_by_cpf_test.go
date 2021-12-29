package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"testing"

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
		wanted  string
		wantErr bool
	}{
		{
			name: "localizada a conta utilizando-se do cpf, e retorna os dados da mesma",
			input: account.Account{
				Name:       "Joao da Silva",
				CPF:        "38330499912",
				Balance:    10000,
			},
			want: account.Account{
				ExternalID: "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:       "Joao da Silva",
				CPF:        "38330499912",
				Balance:    10000,
			},
			wanted:  "38330499912",
			wantErr: false,
		},
		{
			name: "retorna erro ao tentar localizar conta com cpf inexistente",
			input: account.Account{
				Name:       "Joao da Silva",
				CPF:        "38330499912",
				Balance:    10000,
			},
			want:    account.Account{},
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			want: account.Account{
				ID:      "",
				Name:    "",
				CPF:     "",
				Balance: 0,
			},
			wanted:  "38330499999",
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			_, err := accountRepository.Create(ctx, test.input)
			got, err := accountRepository.GetByCPF(ctx, test.wanted)
			if err == nil {
				test.want.CreatedAt = got.CreatedAt
        test.want.ExternalID = got.ExternalID
        test.want.ID = got.ID
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
