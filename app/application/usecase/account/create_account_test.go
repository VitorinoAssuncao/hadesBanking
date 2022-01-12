package account

import (
	"context"
	"database/sql"
	"errors"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	testCases := []struct {
		name        string
		accountMock account.Repository
		tokenMock   token.Repository
		input       account.Account
		want        account.Account
		wantErr     bool
	}{
		{
			name: "conta cadastrada com sucesso, quando dados corretos",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, sql.ErrNoRows
				},
			},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			wantErr: false,
		},
		{
			name: "não é possível cadastrar a conta pois nome está vazio",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, sql.ErrNoRows
				},
			},
			input: account.Account{
				ID:         1,
				Name:       "",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: true,
		},
		{
			name: "não é possível criar a conta, pois cpf já existe",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    0,
					}, nil
				},
			},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: true,
		},
		{
			name: "não é possível criar a conta, pois ocorre um erro na criação da mesma no repositorio",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, nil
				},
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, sql.ErrConnDone
				},
			},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: true,
		},
		{
			name: "não é possível criar a conta, pois cpf já existe",
			accountMock: &account.RepositoryMock{
				CreateFunc: func(ctx context.Context, account account.Account) (account.Account, error) {
					return account, errors.New("test error")
				},
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, sql.ErrNoRows
				},
			},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			want:    account.Account{},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.accountMock, test.tokenMock)
			got, err := u.Create(context.Background(), test.input)
			assert.Equal(t, (err != nil), test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}

}
