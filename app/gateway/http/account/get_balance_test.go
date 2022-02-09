package account

import (
	"context"
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
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/http/middleware"
)

func Test_GetBalance(t *testing.T) {
	t.Parallel()

	const routePattern = "/account/{account_id}/balance"

	type fields struct {
		accountUsecase usecase.Usecase
		authenticator  token.Authenticator
		logger         logHelper.Logger
	}

	type args struct {
		request *http.Request
	}

	testCases := []struct {
		name      string
		fields    fields
		runBefore func(req http.Request)
		args      args
		wantCode  int
		wantBody  map[string]interface{}
	}{
		{
			name: "with a token of login, return the correct value of the account",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						GetBalanceByAccountIDFunc: func(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
							return 0, nil
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{
					ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
						return "c036475f-b7a0-4f34-8f1f-c43515d31724", nil
					},
				},
				logger: &logHelper.RepositoryMock{},
			},
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/account/c036475f-b7a0-4f34-8f1f-c43515d31724/balance", nil),
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"balance": 0,
			},
		},
		{
			name: "without a token of login, return a error",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
							return account.Account{}, nil
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{},
				logger:        &logHelper.RepositoryMock{},
			},
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/account/c036475f-b7a0-4f34-8f1f-c43515d31724/balance", nil),
			},
			wantCode: http.StatusUnauthorized,
			wantBody: map[string]interface{}{
				"error": "authorization token invalid",
			},
		},
		{
			name: "with a token of login, but request for a id that not exist, and return a negative value and error",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						GetBalanceByAccountIDFunc: func(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
							return -1, customError.ErrorAccountIDNotFound
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{
					ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
						return "c036475f-b7a0-4f34-8f1f-c43515d31724", nil
					},
				},
				logger: &logHelper.RepositoryMock{},
			},
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/account/c036475f-b7a0-4f34-8f1f-c43515d31724/balance", nil),
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "account not found, please validate the ID informed",
			},
		},
		{
			name: "with a token of login, but a error happens when trying to find the account in the database, and return a error",
			fields: fields{
				accountUsecase: usecase.New(
					&account.RepositoryMock{
						GetBalanceByAccountIDFunc: func(ctx context.Context, accountID types.ExternalID) (types.Money, error) {
							return -1, customError.ErrorAccountIDSearching
						},
					},
					&token.RepositoryMock{},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{
					ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
						return "c036475f-b7a0-4f34-8f1f-c43515d31724", nil
					},
				},
				logger: &logHelper.RepositoryMock{},
			},
			args: args{
				request: httptest.NewRequest(http.MethodGet, "/account/c036475f-b7a0-4f34-8f1f-c43515d31724/balance", nil),
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: http.StatusInternalServerError,
			wantBody: map[string]interface{}{
				"error": "error when searching for the account",
			},
		},
	}
	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			controller := New(test.fields.accountUsecase, test.fields.authenticator, test.fields.logger)

			req := test.args.request

			if test.runBefore != nil {
				test.runBefore(*req)
			}

			router := mux.NewRouter()
			middleware := middleware.NewMiddleware(test.fields.logger, test.fields.authenticator)
			router.Use(middleware.GetAccountIDFromTokenLogRoutes)
			router.HandleFunc(routePattern, controller.GetBalance).Methods("GET")

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
