package transfer

import (
	"bytes"
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

func Test_Create(t *testing.T) {
	testCases := []struct {
		name            string
		transferUsecase usecase.Usecase
		tokenRepository token.Repository
		logRepository   logHelper.Repository
		input           map[string]interface{}
		runBefore       func(http.Request)
		wantCode        int
		wantBody        map[string]interface{}
	}{
		{
			name: "with the right data, create the transfer sucessfully",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{
							ID:                           1,
							ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
							AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
							AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
							Amount:                       1,
						}, nil
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    100,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"account_destiny_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"amount":             1,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 200,
			wantBody: map[string]interface{}{
				"id":                 "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
				"account_origin_id":  "65d56316-39ad-4937-b41d-be2f103b0bd9",
				"account_destiny_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"value":              0.01,
				"created_at":         "0001-01-01T00:00:00Z",
			},
		},
		{
			name: "try to create a transfer withouth a authorization token, and return a error",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{
							ID:                           1,
							ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
							AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
							AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
							Amount:                       1,
						}, nil
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    100,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "", customError.ErrorServerTokenNotFound
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"account_destiny_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"amount":             1,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "error when generating the authorization token",
			},
		},
		{
			name: "try to create transfer withouth destiny_account_id and return error",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{
							ID:                           1,
							ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
							AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
							AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
							Amount:                       1,
						}, nil
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    100,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"amount": 1,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "field account_origin_id is required",
			},
		},
		{
			name: "try to create a transfer with account origin and destiny equal, and return error",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{
							ID:                           1,
							ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
							AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
							AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
							Amount:                       1,
						}, nil
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    100,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"account_destiny_id": "65d56316-39ad-4937-b41d-be2f103b0bd9",
				"amount":             1,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "origin account and destiny account cannot by the same",
			},
		},
		{
			name: "try to create a transfer with a value of zero, or minor, and return a error",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{
							ID:                           1,
							ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
							AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
							AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
							Amount:                       1,
						}, nil
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    1,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"account_destiny_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"amount":             0,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "transfer amount should be greater than 0",
			},
		},
		{
			name: "try to create a transfer with a value greater than the balance of origin account, and return a error",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{
							ID:                           1,
							ExternalID:                   "cb34f1f3-24ba-4a70-981b-cdc5d77a7347",
							AccountOriginExternalID:      "65d56316-39ad-4937-b41d-be2f103b0bd9",
							AccountDestinationExternalID: "e391600e-7ea9-42aa-85c0-21a2a6c00019",
							Amount:                       1,
						}, nil
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    100,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"account_destiny_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"amount":             101,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 400,
			wantBody: map[string]interface{}{
				"error": "insuficient funds",
			},
		},
		{
			name: "with the right data, create the transfer but has a error in the database, and return a error",
			transferUsecase: usecase.New(
				&transfer.RepositoryMock{
					CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
						return transfer.Transfer{}, customError.ErrorTransferCreate
					}},
				&account.RepositoryMock{
					GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    100,
						}, nil
					}},
				&logHelper.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{
				ExtractAccountIDFromTokenFunc: func(token string) (accountExternalID string, err error) {
					return "65d56316-39ad-4937-b41d-be2f103b0bd9", nil
				}},
			logRepository: &logHelper.RepositoryMock{},
			input: map[string]interface{}{
				"account_destiny_id": "e391600e-7ea9-42aa-85c0-21a2a6c00019",
				"amount":             1,
			},
			runBefore: func(req http.Request) {
				req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI")
			},
			wantCode: 500,
			wantBody: map[string]interface{}{
				"error": "error when creating transfer, please try again",
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			body, _ := json.Marshal(test.input)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewReader(body))
			if test.runBefore != nil {
				test.runBefore(*req)
			}

			controller := New(test.transferUsecase, test.tokenRepository, test.logRepository)
			controller.Create(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != nil {
				wantBodyJson, _ := json.Marshal(test.wantBody)
				assert.JSONEq(t, string(wantBodyJson), rec.Body.String())
			}
		})
	}
}
