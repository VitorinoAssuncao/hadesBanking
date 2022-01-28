package account

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	const routePattern = "/accounts"

	type fields struct {
		accountUsecase usecase.Usecase
		authenticator  token.Authenticator
		logger         logHelper.Logger
	}

	testCases := []struct {
		name     string
		fields   fields
		wantCode int
		wantBody []map[string]interface{}
	}{
		{
			name: "with at last one account existing, data from account is returned sucessfully",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						GetAllFunc: func(ctx context.Context) ([]account.Account, error) {
							return []account.Account{{
								ID:         1,
								Name:       "Joao do Rio",
								ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
								CPF:        "761.647.810-78",
								Secret:     "J0@0doR10",
								Balance:    0,
							}}, nil
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{}},
			wantCode: http.StatusOK,
			wantBody: []map[string]interface{}{{
				"id":         "94b9c27e-2880-42e3-8988-62dceb6b6463",
				"name":       "Joao do Rio",
				"cpf":        "761.647.810-78",
				"balance":    0,
				"created_at": "0001-01-01T00:00:00Z",
			}},
		},
		{
			name: "with one error when listing the accounts, return error for client",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						GetAllFunc: func(ctx context.Context) ([]account.Account, error) {
							return []account.Account{{}}, customError.ErrorAccountsListing
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			wantCode: http.StatusInternalServerError,
			wantBody: []map[string]interface{}{{
				"error": "error when listing all accounts",
			}},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			controller := New(test.fields.accountUsecase, test.fields.authenticator, test.fields.logger)
			req := httptest.NewRequest(http.MethodGet, "/accounts", nil)

			router := mux.NewRouter()
			router.HandleFunc(routePattern, controller.GetAll)

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
