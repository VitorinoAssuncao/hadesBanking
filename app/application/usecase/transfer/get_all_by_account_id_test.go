package transfer

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/transfer"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllByAccountID(t *testing.T) {
	testCases := []struct {
		name         string
		accountMock  account.Repository
		transferMock transfer.Repository
		input        types.ExternalID
		wantQt       int
		wantValue    []transfer.Transfer
		wantErr      error
	}{
		{
			name: "with a valid id, return all transfers made and received by this account",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					account := account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100,
					}
					return account, nil
				},
			},
			transferMock: &transfer.RepositoryMock{
				GetAllByAccountIDFunc: func(ctx context.Context, accountID types.InternalID) ([]transfer.Transfer, error) {
					return []transfer.Transfer{{
						ID:                   1,
						ExternalID:           "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
						AccountOriginID:      1,
						AccountDestinationID: 2,
						Amount:               100,
					}}, nil
				},
			},
			input:  "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
			wantQt: 1,
			wantValue: []transfer.Transfer{{
				ID:                   1,
				ExternalID:           "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
				AccountOriginID:      1,
				AccountDestinationID: 2,
				Amount:               100,
			}},
			wantErr: nil,
		},
		{
			name: "with a invalid id, return error when trying to list all transfers",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{}, customError.ErrorTransferAccountNotFound
				},
			},
			transferMock: &transfer.RepositoryMock{
				GetAllByAccountIDFunc: func(ctx context.Context, accountID types.InternalID) ([]transfer.Transfer, error) {
					return []transfer.Transfer{{
						ID:                   1,
						ExternalID:           "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
						AccountOriginID:      1,
						AccountDestinationID: 2,
						Amount:               100,
					}}, nil
				},
			},
			input:     "29e5ff6b-6a0d-402e-a77c-67fc2281aca9",
			wantQt:    0,
			wantValue: []transfer.Transfer{},
			wantErr:   customError.ErrorTransferAccountNotFound,
		},
		{
			name: "when listing all transfers, the database return a error",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100,
					}, nil
				},
			},
			transferMock: &transfer.RepositoryMock{
				GetAllByAccountIDFunc: func(ctx context.Context, accountID types.InternalID) ([]transfer.Transfer, error) {
					return []transfer.Transfer{}, customError.ErrorTransferListing
				},
			},
			input:     "29e5ff6b-6a0d-402e-a77c-67fc2281aca0",
			wantQt:    0,
			wantValue: []transfer.Transfer{},
			wantErr:   customError.ErrorTransferListing,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.transferMock, test.accountMock)
			got, err := u.GetAllByAccountID(context.Background(), test.input)
			assert.Equal(t, err, test.wantErr)
			assert.Equal(t, test.wantQt, len(got))
			assert.Equal(t, test.wantValue, got)
		})
	}
}
