package account

import (
	"context"
	"database/sql"
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
				GetBalanceByAccountIDFunc: func(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
					return types.Money(100), nil
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
				GetBalanceByAccountIDFunc: func(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
					return -1, customError.ErrorAccountIDNotFound
				},
			},
			logMock: &logHelper.RepositoryMock{},
			input:   "94b9c27e-2880-42e3-8988-62dceb6b6464",
			want:    -1,
			wantErr: customError.ErrorAccountIDNotFound,
		},
		{
			name: "with the correct id, have a sudenly error in database and return a error",
			accountMock: &account.RepositoryMock{
				GetBalanceByAccountIDFunc: func(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
					return -1, sql.ErrConnDone
				},
			},
			logMock: &logHelper.RepositoryMock{},
			input:   "94b9c27e-2880-42e3-8988-62dceb6b6464",
			want:    -1,
			wantErr: customError.ErrorAccountIDSearching,
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
