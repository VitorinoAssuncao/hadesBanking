package transfer

import (
	"context"
	"net/http"
	"net/http/httptest"
	usecase "stoneBanking/app/application/usecase/transfer"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/http/middleware"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllByAccountID(t *testing.T) {
	const routePattern = "/transfers"

	type fields struct {
		transferUsecase usecase.Usecase
		authenticator   token.Authenticator
		logger          logHelper.Logger
	}

	testCases := []struct {
		name      string
		fields    fields
		runBefore func(http.Request)
		wantCode  int
		wantBody  string
	}{
		{
			name: "with transfers in the database, return the data successfully",
			fields: fields{
				transferUsecase: usecase.New(
					&transfer.RepositoryMock{
						GetAllByAccountIDFunc: func(ctx context.Context, accountID types.InternalID) ([]transfer.Transfer, error) {
							return []transfer.Transfer{{
								ID:                           1,
								ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
								AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
								AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
								Amount:                       1,
							}}, nil
						}},
					&account.RepositoryMock{
						GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
							return account.Account{}, nil
						}},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{
					ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
						return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
					}},
				logger: &logHelper.RepositoryMock{},
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: http.StatusOK,
			wantBody: `
				[
					{
						"id":                     "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
						"account_origin_id":      "65d56316-39ad-4937-b41d-be2f103b0bd9",
						"account_destination_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
						"value":                  0.01,
						"created_at":             "0001-01-01T00:00:00Z"
					}
				]`,
		},
		{
			name: "try to get all transfers without a token and return error",
			fields: fields{
				transferUsecase: usecase.New(
					&transfer.RepositoryMock{
						GetAllByAccountIDFunc: func(ctx context.Context, accountID types.InternalID) ([]transfer.Transfer, error) {
							return []transfer.Transfer{{}}, nil
						}},
					&account.RepositoryMock{
						GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
							return account.Account{}, nil
						}},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{
					ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
						return "", customError.ErrorServerTokenNotFound
					}},
				logger: &logHelper.RepositoryMock{},
			},
			wantCode: http.StatusUnauthorized,
			wantBody: `{"error": "authorization token invalid"}`,
		},
		{
			name: "try to get all transfers, but a error in database happens, and return a error to user",
			fields: fields{
				transferUsecase: usecase.New(
					&transfer.RepositoryMock{
						GetAllByAccountIDFunc: func(ctx context.Context, accountID types.InternalID) ([]transfer.Transfer, error) {
							return []transfer.Transfer{{}}, customError.ErrorTransferListing
						}},
					&account.RepositoryMock{
						GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
							return account.Account{}, nil
						}},
					&logHelper.RepositoryMock{}),
				authenticator: &token.RepositoryMock{
					ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
						return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
					}},
				logger: &logHelper.RepositoryMock{},
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: http.StatusInternalServerError,
			wantBody: `{"error": "error when listing all transfers"}`,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/transfers", nil)
			if test.runBefore != nil {
				test.runBefore(*req)
			}

			router := mux.NewRouter()

			middleware := middleware.NewMiddleware(test.fields.logger, test.fields.authenticator)
			router.Use(middleware.GetAccountIDFromTokenLogRoutes)

			controller := New(test.fields.transferUsecase, test.fields.authenticator, test.fields.logger)
			router.HandleFunc(routePattern, controller.GetAllByAccountID)

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			assert.JSONEq(t, test.wantBody, rec.Body.String())
		})
	}
}
