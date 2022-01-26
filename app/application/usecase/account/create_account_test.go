package account

import (
	"context"
<<<<<<< HEAD
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

=======
>>>>>>> refactor: corrected test that failed as sideeffect
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
)

func Test_Create(t *testing.T) {
	testCases := []struct {
		name        string
		accountMock account.Repository
		tokenMock   token.Authenticator
		logMock     logHelper.Logger
		input       account.Account
		want        account.Account
		wantErr     error
	}{
		{
			name: "with the right data, create account successfully",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
<<<<<<< HEAD
				GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, customError.ErrorAccountCPFNotFound
				},
=======
>>>>>>> refactor: corrected test that failed as sideeffect
			},
			logMock: &logHelper.RepositoryMock{},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			wantErr: nil,
		},
		{
			name: "with data missing the Name, as not possible to create account, and return a error",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
<<<<<<< HEAD
				GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, customError.ErrorAccountCPFNotFound
				},
=======
>>>>>>> refactor: corrected test that failed as sideeffect
			},
			logMock: &logHelper.RepositoryMock{},
			input: account.Account{
				ID:         1,
				Name:       "",
				ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: customError.ErrorAccountNameRequired,
		},
		{
			name: "with right input data, try to create a account, but is duplicated from one that exist, and return error",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
<<<<<<< HEAD
					return account, nil
				},
				GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    0,
					}, nil
=======
					return account, customError.ErrorAccountCPFExists
>>>>>>> refactor: corrected test that failed as sideeffect
				},
			},
			logMock: &logHelper.RepositoryMock{},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: customError.ErrorAccountCPFExists,
		},
		{
			name: "with right data, try to create a account, but has a error when validating if a account with that cpf exist",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
				GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, sql.ErrConnDone
				},
			},
			logMock: &logHelper.RepositoryMock{},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: customError.ErrorCreateAccount,
		},
		{
			name: "with the right data, try to create a account but has a error when creating in the database",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, customError.ErrorCreateAccount
				},
				GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, customError.ErrorAccountCPFNotFound
				},
			},
			logMock: &logHelper.RepositoryMock{},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: customError.ErrorCreateAccount,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.accountMock, test.tokenMock, test.logMock)
			got, err := u.Create(context.Background(), test.input)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}

}
