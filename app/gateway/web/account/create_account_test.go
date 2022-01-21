package account

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/web/account/vo/input"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	testCases := []struct {
		name            string
		accountUsecase  usecase.Usecase
		tokenRepository token.Repository
		logRepository   logHelper.Repository
		input           input.CreateAccountVO
		wantCode        int
		wantBody        map[string]interface{}
	}{
		{
			name: "with the right data, account is created sucessfully",
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
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{}, customError.ErrorAccountCPFNotFound
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{},
			logRepository:   &logHelper.RepositoryMock{},
			input: input.CreateAccountVO{
				Name:    "Joao",
				CPF:     "761.647.810-78",
				Secret:  "J0@0doR10",
				Balance: 100,
			},
			wantCode: 200,
			wantBody: map[string]interface{}{
				"id":         "94b9c27e-2880-42e3-8988-62dceb6b6463",
				"name":       "Joao do Rio",
				"cpf":        "761.647.810-78",
				"balance":    0,
				"created_at": "0001-01-01T00:00:00Z",
			},
		},
		{
			name: "data from input withouth name, generating error in validation",
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
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{}, sql.ErrNoRows
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{},
			logRepository:   &logHelper.RepositoryMock{},
			input: input.CreateAccountVO{
				Name:    "",
				CPF:     "761.647.810-78",
				Secret:  "J0@0doR10",
				Balance: 100,
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "the field 'Name' is required",
			},
		},
		{
			name: "data from input with a negative ammount in origin, generating error in validation",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
						return account.Account{}, nil
					},
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{}, sql.ErrNoRows
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{},
			logRepository:   &logHelper.RepositoryMock{},
			input: input.CreateAccountVO{
				Name:    "Joao do Rio",
				CPF:     "761.647.810-78",
				Secret:  "J0@0doR10",
				Balance: -100,
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "the field 'Balance' need to by equal or major than 0(zero)",
			},
		},
		{
			name: "with correct data, but have a error when creating the account in database",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
						return account.Account{}, customError.ErrorCreateAccount
					},
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{}, sql.ErrNoRows
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{},
			logRepository:   &logHelper.RepositoryMock{},
			input: input.CreateAccountVO{
				Name:    "Joao do Rio",
				CPF:     "761.647.810-78",
				Secret:  "J0@0doR10",
				Balance: 100,
			},
			wantCode: 500,
			wantBody: map[string]interface{}{
				"error": "error when creating a new account",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			body, _ := json.Marshal(test.input)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			controller := New(test.accountUsecase, test.tokenRepository, test.logRepository)
			controller.Create(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != nil {
				wantBodyJson, _ := json.Marshal(test.wantBody)
				assert.JSONEq(t, string(wantBodyJson), rec.Body.String())
			}
		})
	}

}
