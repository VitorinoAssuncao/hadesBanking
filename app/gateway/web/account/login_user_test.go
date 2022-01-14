package account

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/account/vo/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoginUser(t *testing.T) {
	testCases := []struct {
		name            string
		accountUsecase  usecase.Usecase
		tokenRepository token.Repository
		input           input.CreateAccountVO
		wantCode        int
		wantBody        map[string]interface{}
	}{
		{
			name: "with the correct data, log the user and return the authorization token",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							Balance:    0,
						}, nil
					},
				},
				&token.RepositoryMock{
					GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
						signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
						return signedToken, nil
					},
				}),
			tokenRepository: &token.RepositoryMock{},
			input: input.CreateAccountVO{
				CPF:    "761.647.810-78",
				Secret: "12344",
			},
			wantCode: 200,
			wantBody: map[string]interface{}{
				"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
			},
		},
		{
			name: "with the login data withouth cpf, when validating the data return a error",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							Balance:    0,
						}, nil
					},
				},
				&token.RepositoryMock{
					GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
						signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
						return signedToken, nil
					},
				}),
			tokenRepository: &token.RepositoryMock{},
			input: input.CreateAccountVO{
				CPF:    "",
				Secret: "12344",
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": customError.ErrorAccountLogin.Error(),
			},
		},
		{
			name: "with the correct cpf but the wrong password, try to log, but return a error",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							Balance:    0,
						}, nil
					},
				},
				&token.RepositoryMock{
					GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
						signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
						return signedToken, nil
					},
				}),
			tokenRepository: &token.RepositoryMock{},
			input: input.CreateAccountVO{
				CPF:    "761.647.810-78",
				Secret: "12345",
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": customError.ErrorAccountLogin.Error(),
			},
		},
		{
			name: "with the correct data, but happens a error when generating the token, and return error",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							Balance:    0,
						}, nil
					},
				},
				&token.RepositoryMock{
					GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
						return "", customError.ErrorAccountTokenGeneration
					},
				}),
			tokenRepository: &token.RepositoryMock{},
			input: input.CreateAccountVO{
				CPF:    "761.647.810-78",
				Secret: "12344",
			},
			wantCode: 500,
			wantBody: map[string]interface{}{
				"error": customError.ErrorAccountTokenGeneration.Error(),
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.input)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/account/login", bytes.NewReader(body))
			controller := New(test.accountUsecase, test.tokenRepository)
			controller.LoginUser(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != nil {
				wantBodyJson, _ := json.Marshal(test.wantBody)
				assert.JSONEq(t, string(wantBodyJson), rec.Body.String())
			}
		})
	}
}
