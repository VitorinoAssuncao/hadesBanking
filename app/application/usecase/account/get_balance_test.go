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
						ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100,
					}
					return account, nil
				},
			},
			input:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
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
			input:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeXXX",
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
