package account

import (
	"context"
	"database/sql"
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
		name      string
		input     account.Account
		runBefore func(db *sql.DB)
		want      account.Account
		wantErr   bool
	}{
		{
			name: "conta cadastrada com sucesso, quando dados corretos",
			input: account.Account{
				ID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
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
			name: "ao tentar criar a conta apresenta que já existe conta cadastrada com este cpf",
			input: account.Account{
				ID:      "d3280f8c-570a-450d-89f7-3509bc84980d",
				Name:    "Joao da Silva",
				CPF:     "38330499912",
				Balance: 10000,
			},
			runBefore: func(db *sql.DB) {
				sqlQuery :=
					`
				INSERT INTO
					accounts (id, name, cpf, secret, balance, created_at)
				VALUES
					('d3280f8c-570a-450d-89f7-3509bc84980d', 'Joao da Silva', '38330499912', 'password', 100, $1)
				`
				_, err := db.Exec(sqlQuery, time.Now())
				if err != nil {
					t.Errorf(err.Error())
				}
			},
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
			TruncateTable(database)
			if test.runBefore != nil {
				test.runBefore(database)
			}
			got, err := accountRepository.Create(ctx, test.input)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
