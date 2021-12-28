package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	now := time.Now()
	testCases := []struct {
		name    string
		input   account.Account
		want    account.Account
		wanted  string
		wantErr bool
	}{
		{
			name: "localizado a conta usando o ID externo (uuid), e retorna os dados da mesma ",
			input: account.Account{
				ID:        "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: now,
			},
			want: account.Account{
				ID:        "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: now,
			},
			wanted:  "d3280f8c-570a-450d-89f7-3509bc84980d",
			wantErr: false,
		}, {
			name: "retorna dados vazios e erro, ao tentar localizar conta com ID inexistente",
			input: account.Account{
				ID:        "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:      "Joao da Silva",
				CPF:       "38330499912",
				Balance:   10000,
				CreatedAt: now,
			},
			wanted: "d3280f8c-570a-450d-89f7-3509bc849899",
			want: account.Account{
				ID:      "",
				Name:    "",
				CPF:     "",
				Balance: 0,
			},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := accountRepository.Create(ctx, test.input)
			got, err := accountRepository.GetByID(ctx, types.AccountID(test.wanted))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
