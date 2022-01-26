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
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBalance(t *testing.T) {
	testCases := []struct {
		name           string
		accountUsecase usecase.Usecase
		authenticator  token.Authenticator
		logger         logHelper.Logger
		runBefore      func(req http.Request)
		wantCode       int
		wantBody       map[string]interface{}
	}{
		{
			name: "with a token of login, return the correct value of the account",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    0,
						}, nil
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			authenticator: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "94b9c27e-2880-42e3-8988-62dceb6b6463", nil
				},
			},
			logger: &logHelper.RepositoryMock{},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 200,
			wantBody: map[string]interface{}{
				"balance": 0,
			},
		},
		{
			name: "without a token of login, return a error",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{}, nil
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			authenticator: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "", customError.ErrorServerTokenNotFound
				},
			},
			logger:   &logHelper.RepositoryMock{},
			wantCode: 401,
			wantBody: map[string]interface{}{
				"error": "authorization token invalid",
			},
		},
		{
			name: "with a token of login, but a error happens when trying to find the account in the database, and return a error",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{}, customError.ErrorAccountIDNotFound
					},
				},
				&token.RepositoryMock{},
				&logHelper.RepositoryMock{}),
			authenticator: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "94b9c27e-2880-42e3-8988-62dceb6b6463", nil
				},
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			logger:   &logHelper.RepositoryMock{},
			wantCode: 500,
			wantBody: map[string]interface{}{
				"error": "account not found, please validate the ID informed",
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/account/balance", nil)

			if test.runBefore != nil {
				test.runBefore(*req)
			}

			controller := New(test.accountUsecase, test.authenticator, test.logger)
			controller.GetBalance(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != nil {
				wantBodyJson, _ := json.Marshal(test.wantBody)
				assert.JSONEq(t, string(wantBodyJson), rec.Body.String())
			}
		})
	}
}
