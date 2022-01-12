package account

import (
	"context"
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	testCases := []struct {
		name        string
		accountMock account.Repository
		tokenMock   token.Repository
		runBefore   func(db *sql.DB)
		want        int
		wantErr     bool
	}{
		{
			name: "retorna uma conta cadastrada com sucesso",
			accountMock: &account.RepositoryMock{
				GetAllFunc: func(ctx context.Context) ([]account.Account, error) {
					tempAccount := account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    0,
					}
					var accounts = make([]account.Account, 0)
					accounts = append(accounts, tempAccount)
					return accounts, nil
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "erro ao tentar buscar as contas no banco",
			accountMock: &account.RepositoryMock{
				GetAllFunc: func(ctx context.Context) ([]account.Account, error) {
					return []account.Account{}, customError.ErrorAccountsListing
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.accountMock, test.tokenMock)
			got, err := u.GetAll(context.Background())
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
