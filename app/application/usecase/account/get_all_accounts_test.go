package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
)

func Test_GetAll(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		accountMock account.Repository
		tokenMock   token.Authenticator
		logMock     logHelper.Logger
		want        int
		wantErr     error
	}{
		{
			name: "return all accounts that exist in the database",
			accountMock: &account.RepositoryMock{
				GetAllFunc: func(ctx context.Context) ([]account.Account, error) {
					tempAccount := account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    0,
					}
					var accounts = make([]account.Account, 0)
					accounts = append(accounts, tempAccount)
					return accounts, nil
				},
			},
			logMock: &logHelper.RepositoryMock{},
			want:    1,
			wantErr: nil,
		},
		{
			name: "when trying to list all the accounts in the database, as a error in the query",
			accountMock: &account.RepositoryMock{
				GetAllFunc: func(ctx context.Context) ([]account.Account, error) {
					return []account.Account{}, customError.ErrorAccountsListing
				},
			},
			logMock: &logHelper.RepositoryMock{},
			want:    0,
			wantErr: customError.ErrorAccountsListing,
		},
	}
	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			u := New(test.accountMock, test.tokenMock, test.logMock)
			got, err := u.GetAll(context.Background())
			assert.Equal(t, err, test.wantErr)
			assert.Equal(t, test.want, len(got))
		})
	}
}
