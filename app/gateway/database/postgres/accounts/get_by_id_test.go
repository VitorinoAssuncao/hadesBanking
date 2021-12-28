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
	now := time.Now()
	testCases := []struct {
		name    string
		input   account.Account
		want    account.Account
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
			want: account.Account{
				ID:      "d3280f8c-570a-450d-89f7-3509bc849899",
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
			got, err := accountRepository.GetByID(ctx, test.want.ID)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want.Name, got.Name)
		})
	}
}
