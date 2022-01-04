package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateBalance(t *testing.T) {
	ctx := context.Background()
	database := databaseTest
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name       string
		want       bool
		runBefore  func() (value types.AccountExternalID)
		inputID    string
		inputValue int
		wantErr    bool
	}{
		{
			name: "faz a atualização do saldo com sucesso, em uma conta que existe",
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
			want:       true,
			inputValue: 100,
			wantErr:    false,
		},
		{
			name:       "gera erro, pois ao tentar atualizar o saldo, a conta de destino não existe",
			want:       false,
			inputID:    "",
			inputValue: 100,
			wantErr:    true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			if test.runBefore != nil {
				test.inputID = string(test.runBefore())
			}
			got, err := accountRepository.UpdateBalance(ctx, test.inputValue, types.AccountExternalID(test.inputID))
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
