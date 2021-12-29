package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"testing"

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
			name: "localiza todas as contas, quando existe ao menos uma cadastrada",
			input: account.Account{
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			TruncateTable(database)
			_, err := accountRepository.Create(ctx, test.input)
			got, err := accountRepository.GetAll(ctx)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
