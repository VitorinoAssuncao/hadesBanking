package account

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/account/vo/input"
)

func Test_Create(t *testing.T) {
	const routePattern = "/accounts"

	type fields struct {
		accountUsecase usecase.Usecase
		authenticator  token.Authenticator
		logger         logHelper.Logger
	}

	type args struct {
		input input.CreateAccountVO
	}

	testCases := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody string
	}{
		{
			name: "with the right data, account is created successfully",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
							return account.Account{
								ID:         1,
								Name:       "Joao do Rio",
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "J0@0doR10",
								Balance:    0,
							}, nil
						},
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{}, customError.ErrorAccountCPFNotFound
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.CreateAccountVO{
					Name:    "Joao",
					CPF:     "761.647.810-78",
					Secret:  "J0@0doR10",
					Balance: 100,
				},
			},
			wantCode: http.StatusCreated,
			wantBody: `
			{
				"id":         "94b9c27e-2880-42e3-8988-62dceb6b6463",
				"name":       "Joao do Rio",
				"cpf":        "761.647.810-78",
				"balance":    0,
				"created_at": "0001-01-01T00:00:00Z"
			}`,
		},
		{
			name: "data from input without name, generating error in validation",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
							return account.Account{
								ID:         1,
								Name:       "Joao do Rio",
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "J0@0doR10",
								Balance:    0,
							}, nil
						},
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{}, sql.ErrNoRows
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.CreateAccountVO{
					Name:    "",
					CPF:     "761.647.810-78",
					Secret:  "J0@0doR10",
					Balance: 100,
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: `[
				{
					"error": "the field 'Name' is required"
				}
			]`,
		},
		{
			name: "data from input with a negative amount in origin, generating error in validation",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
							return account.Account{}, nil
						},
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{}, sql.ErrNoRows
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.CreateAccountVO{
					Name:    "Joao do Rio",
					CPF:     "761.647.810-78",
					Secret:  "J0@0doR10",
					Balance: -100,
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: `[
				{
					"error": "the field 'Balance' need to by equal or major than 0(zero)"
				}
				]`,
		},
		{
			name: "with correct data, but have a error when creating the account in database",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
							return account.Account{}, customError.ErrorCreateAccount
						},
						GetCredentialByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
							return account.Account{}, sql.ErrNoRows
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				input: input.CreateAccountVO{
					Name:    "Joao do Rio",
					CPF:     "761.647.810-78",
					Secret:  "J0@0doR10",
					Balance: 100,
				},
			},
			wantCode: http.StatusInternalServerError,
			wantBody: `{
				"error": "error when creating a new account"
			}`,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			controller := New(test.fields.accountUsecase, test.fields.authenticator, test.fields.logger)

			body, _ := json.Marshal(test.args.input)
			req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))

			router := mux.NewRouter()
			router.HandleFunc(routePattern, controller.Create)

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			assert.JSONEq(t, test.wantBody, rec.Body.String())
		})
	}

}
