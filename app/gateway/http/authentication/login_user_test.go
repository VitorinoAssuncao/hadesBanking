package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	usecase "stoneBanking/app/application/usecase/authentication"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/authentication/vo/input"
)

func Test_LoginUser(t *testing.T) {
	const routePattern = "/login"

	type fields struct {
		authUsecase   usecase.Usecase
		authenticator token.Authenticator
		logger        logHelper.Logger
	}

	type args struct {
		input input.LoginVO
	}

	testCases := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody map[string]interface{}
	}{
		{
			name: "with the correct data, log the user and return the authorization token",
			fields: fields{
				authUsecase: usecase.New(
					&account.RepositoryMock{
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							}, nil
						},
					},
					&token.RepositoryMock{
						GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
							signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
							return signedToken, nil
						},
					},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.LoginVO{
					CPF:    "761.647.810-78",
					Secret: "12344",
				},
			},
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
			},
		},
		{
			name: "with the login data without cpf, when validating the data return a error",
			fields: fields{
				authUsecase: usecase.New(
					&account.RepositoryMock{
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							}, nil
						},
					},
					&token.RepositoryMock{
						GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
							signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
							return signedToken, nil
						},
					},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.LoginVO{
					CPF:    "",
					Secret: "12344",
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "cpf or secret invalid, please validate then",
			},
		},
		{
			name: "with the correct cpf but the wrong password, try to log, but return a error",
			fields: fields{
				authUsecase: usecase.New(
					&account.RepositoryMock{
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							}, nil
						},
					},
					&token.RepositoryMock{
						GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
							signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
							return signedToken, nil
						},
					},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.LoginVO{
					CPF:    "761.647.810-78",
					Secret: "12345",
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "cpf or secret invalid, please validate then",
			},
		},
		{
			name: "with the correct data, but happens a error when generating the token, and return error",
			fields: fields{
				authUsecase: usecase.New(
					&account.RepositoryMock{
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "$2a$14$zmq6uXNuf.1ZwUHHDDAnR.ggIYn0YEnWF/1HeeZrf8d8B55mkk.aq",
							}, nil
						},
					},
					&token.RepositoryMock{
						GenerateTokenFunc: func(accountExternalID string) (signedToken string, err error) {
							return "", customError.ErrorAccountTokenGeneration
						},
					},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.LoginVO{
					CPF:    "761.647.810-78",
					Secret: "12344",
				},
			},
			wantCode: http.StatusInternalServerError,
			wantBody: map[string]interface{}{
				"error": "error when generating the authorization token",
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			controller := New(test.fields.authUsecase, test.fields.authenticator, test.fields.logger)

			body, _ := json.Marshal(test.args.input)
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))

			router := mux.NewRouter()
			router.HandleFunc(routePattern, controller.LoginUser)

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != nil {
				wantBodyJson, _ := json.Marshal(test.wantBody)
				assert.JSONEq(t, string(wantBodyJson), rec.Body.String())
			}
		})
	}
}
