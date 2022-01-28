package transfer

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	usecase "stoneBanking/app/application/usecase/transfer"
	"stoneBanking/app/domain/entities/account"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllByAccountID(t *testing.T) {
	testCases := []struct {
		name            string
		transferUsecase usecase.Usecase
		authenticator   token.Authenticator
		logger          logHelper.Logger
		runBefore       func(http.Request)
		wantCode        int
		wantBody        []map[string]interface{}
	}{
		{
			name: "with transfers in the database, return the data sucessfully",
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
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: http.StatusOK,
			wantBody: []map[string]interface{}{{
				"id":                     "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
				"account_origin_id":      "65d56316-39ad-4937-b41d-be2f103b0bd9",
				"account_destination_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"value":                  0.01,
				"created_at":             "0001-01-01T00:00:00Z",
			}},
		},
		{
			name: "try to get all transfers withouth a token and return error",
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
			logger: &logHelper.RepositoryMock{},
			authenticator: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "", customError.ErrorServerTokenNotFound
				}},
			wantCode: http.StatusBadRequest,
			wantBody: []map[string]interface{}{{
				"error": "authorization token invalid",
			}},
		},
		{
			name: "try to get all transfers, but a error in database happens, and return a error to user",
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
			logger: &logHelper.RepositoryMock{},
			authenticator: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: http.StatusInternalServerError,
			wantBody: []map[string]interface{}{{
				"error": "error when listing all transfers",
			}},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/transfer", nil)
			if test.runBefore != nil {
				test.runBefore(*req)
			}

			controller := New(test.transferUsecase, test.authenticator, test.logger)
			controller.GetAllByAccountID(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != nil {
				wantBodyJson, _ := json.Marshal(test.wantBody)
				assert.JSONEq(t, string(wantBodyJson), rec.Body.String())
			}
		})
	}
}
