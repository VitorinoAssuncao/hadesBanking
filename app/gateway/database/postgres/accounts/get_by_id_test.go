package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetByID(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name      string
		want      account.Account
		runBefore func() (value types.AccountExternalID)
		input     types.AccountExternalID
		wantErr   bool
	}{
		{
			name: "localizado a conta usando o ID externo (uuid), e retorna os dados da mesma ",
			runBefore: func() (value types.AccountExternalID) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Balance: 10000,
				}
				created, err := accountRepository.Create(ctx, input)

				if err == nil {
					value = created.ExternalID
				}

				return value
			},
			want: account.Account{
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			input:   "d3280f8c-570a-450d-89f7-3509bc84980d",
			wantErr: false,
		}, {
			name:    "retorna dados vazios e erro, ao tentar localizar conta com ID inexistente",
			want:    account.Account{},
			input:   "d3280f8c-570a-450d-89f7-3509bc849899",
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			if test.runBefore != nil {
				test.input = test.runBefore()
			}
			got, err := accountRepository.GetByID(ctx, types.AccountExternalID(test.input))
			if err == nil {
				test.want.CreatedAt = got.CreatedAt
				test.want.ID = got.ID
				test.want.ExternalID = got.ExternalID
			}
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
