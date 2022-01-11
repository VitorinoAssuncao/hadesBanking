package account

import (
	"context"
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (usecase *usecase) Test_Create(t *testing.T) {
	testCases := []struct {
		name      string
		input     account.RepositoryMock
		runBefore func(db *sql.DB)
		want      account.Account
		wantErr   bool
	}{
		{
			name: "conta cadastrada com sucesso, quando dados corretos",
			input: account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
			},
			want:    account.Account{},
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			usecase := New(test.input, "")

			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
