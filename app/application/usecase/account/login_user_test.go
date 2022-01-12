package account

import (
	"context"
	"errors"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoginUser(t *testing.T) {
	testCases := []struct {
		name        string
		accountMock account.Repository
		tokenMock   token.Repository
		input       account.Account
		want        string
		wantErr     error
	}{
		{
			name: "dados que dados de login estejam corretos, retorna token de autenticação",
			accountMock: &account.RepositoryMock{
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "4873e8ff-5f46-417b-930f-f3d914a19df2",
						CPF:        "761.647.810-78",
						Secret:     types.Password("J0@0doR10").Hash(),
						Balance:    0,
					}, nil
				},
			},
			tokenMock: &token.RepositoryMock{
				GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
					signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
					return signedToken, nil
				},
			},
			input: account.Account{
				CPF:    "761.647.810-78",
				Secret: "J0@0doR10",
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
			wantErr: nil,
		},
		{
			name: "retornará com erro ao informar conta incorretamente",
			accountMock: &account.RepositoryMock{
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{}, errors.New("test error")
				},
			},
			tokenMock: &token.RepositoryMock{
				GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
					signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
					return signedToken, nil
				},
			},
			input: account.Account{
				CPF:    "761.647.810-78",
				Secret: "J0@0doR10",
			},
			want:    "",
			wantErr: customError.ErrorAccountLogin,
		},
		{
			name: "dados senha incorreta, deve retornar sem o token e apresentando erro",
			accountMock: &account.RepositoryMock{
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "4873e8ff-5f46-417b-930f-f3d914a19df2",
						CPF:        "761.647.810-78",
						Secret:     types.Password("J0@0doR11").Hash(),
						Balance:    0,
					}, nil
				},
			},
			tokenMock: &token.RepositoryMock{
				GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
					signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
					return signedToken, nil
				},
			},
			input: account.Account{
				CPF:    "761.647.810-78",
				Secret: "J0@0doR10",
			},
			want:    "",
			wantErr: customError.ErrorAccountLogin,
		},
		{
			name: "deverá apresentar ao erro, ao ocorrer falha na geração do token",
			accountMock: &account.RepositoryMock{
				GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "4873e8ff-5f46-417b-930f-f3d914a19df2",
						CPF:        "761.647.810-78",
						Secret:     types.Password("J0@0doR10").Hash(),
						Balance:    0,
					}, nil
				},
			},
			tokenMock: &token.RepositoryMock{
				GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
					return signedToken, errors.New("test error in token generation")
				},
			},
			input: account.Account{
				CPF:    "761.647.810-78",
				Secret: "J0@0doR10",
			},
			want:    "",
			wantErr: customError.ErrorAccountTokenGeneration,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.accountMock, test.tokenMock)
			got, err := u.LoginUser(context.Background(), test.input)
			assert.Equal(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
