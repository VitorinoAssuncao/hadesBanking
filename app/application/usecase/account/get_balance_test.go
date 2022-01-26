package account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
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
		tokenMock   token.Authenticator
		logMock     logHelper.Logger
		input       string
		want        float64
		wantErr     error
	}{
		{
			name: "with the right id, return the balance of account",
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
			logMock: &logHelper.RepositoryMock{},
			input:   "94b9c27e-2880-42e3-8988-62dceb6b6463",
			want:    1,
			wantErr: nil,
		},
		{
			name: "with a invalid id, return a error of account not found",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{}, customError.ErrorAccountIDNotFound
				},
			},
			logMock: &logHelper.RepositoryMock{},
			input:   "94b9c27e-2880-42e3-8988-62dceb6b6464",
			want:    -1,
			wantErr: customError.ErrorAccountIDNotFound,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.accountMock, test.tokenMock, test.logMock)
			got, err := u.GetBalance(context.Background(), test.input)
			assert.Equal(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
