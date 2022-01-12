package account

import (
	"context"
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBalance(t *testing.T) {
	testCases := []struct {
		name        string
		accountMock account.Repository
		tokenMock   token.Repository
		input       string
		runBefore   func(db *sql.DB)
		want        float64
		wantErr     error
	}{
		{
			name: "dado id externo correto, retorna o valor da conta apropriadamente",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					account := account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100,
					}
					return account, nil
				},
			},
			input:   "94b9c27e-2880-42e3-8988-62dceb6b6463",
			want:    1,
			wantErr: nil,
		},
		{
			name: "dado id incorreto, retorna valor negativo e erro",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{}, customError.ErrorAccountIDNotFound
				},
			},
			input:   "94b9c27e-2880-42e3-8988-62dceb6b6464",
			want:    -1,
			wantErr: customError.ErrorAccountIDNotFound,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.accountMock, test.tokenMock)
			got, err := u.GetBalance(context.Background(), test.input)
			assert.Equal(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
