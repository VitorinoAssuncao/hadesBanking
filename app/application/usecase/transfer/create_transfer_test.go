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

func Test_Create(t *testing.T) {
	testCases := []struct {
		name         string
		accountMock  account.Repository
		transferMock transfer.Repository
		input        transfer.Transfer
		want         transfer.Transfer
		wantErr      error
	}{
		{
			name: "with the correct data, sucessfully create a transfer and update the accounts",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100}, nil
				},
			},
			transferMock: &transfer.RepositoryMock{
				CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
					return transfer.Transfer{
						ID:                           1,
						ExternalID:                   "56286ebe-8798-40ba-81aa-3caa74197cd1",
						AccountOriginID:              1,
						AccountOriginExternalID:      "01aacb75-cbd4-45a9-91ed-6cf2f6dcf772",
						AccountDestinationID:         2,
						AccountDestinationExternalID: "f53420f2-616c-4fe3-a957-84f03386a82f",
						Amount:                       1,
					}, nil
				},
			},
			input: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 2,
				Amount:               1,
			},
			want: transfer.Transfer{
				ID:                           1,
				ExternalID:                   "56286ebe-8798-40ba-81aa-3caa74197cd1",
				AccountOriginID:              1,
				AccountOriginExternalID:      "01aacb75-cbd4-45a9-91ed-6cf2f6dcf772",
				AccountDestinationID:         2,
				AccountDestinationExternalID: "f53420f2-616c-4fe3-a957-84f03386a82f",
				Amount:                       1,
			},
			wantErr: nil,
		},
		{
			name: "with wrong origin account, return error",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{}, customError.ErrorTransferCreateOriginError
				},
			},
			transferMock: &transfer.RepositoryMock{
				CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
					return transfer.Transfer{}, nil
				},
			},
			input: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 2,
				Amount:               1,
			},
			want:    transfer.Transfer{},
			wantErr: customError.ErrorTransferCreateOriginError,
		},
		{
			name: "with right data, but the fund is insufficient, and return a error",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100}, nil
				},
			},
			transferMock: &transfer.RepositoryMock{
				CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
					return transfer.Transfer{}, customError.ErrorTransferCreateInsufficientFunds
				},
			},
			input: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 2,
				Amount:               101,
			},
			want:    transfer.Transfer{},
			wantErr: customError.ErrorTransferCreateInsufficientFunds,
		},
		{
			name: "with right data, have a error in database",
			accountMock: &account.RepositoryMock{
				GetByIDFunc: func(ctx context.Context, accountID types.ExternalID) (account.Account, error) {
					return account.Account{
						ID:         1,
						Name:       "Joao do Rio",
						ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
						CPF:        "761.647.810-78",
						Secret:     "J0@0doR10",
						Balance:    100}, nil
				},
			},
			transferMock: &transfer.RepositoryMock{
				CreateFunc: func(ctx context.Context, transferData transfer.Transfer) (transfer.Transfer, error) {
					return transfer.Transfer{}, customError.ErrorTransferCreate
				},
			},
			input: transfer.Transfer{
				AccountOriginID:      1,
				AccountDestinationID: 2,
				Amount:               1,
			},
			want:    transfer.Transfer{},
			wantErr: customError.ErrorTransferCreate,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			u := New(test.transferMock, test.accountMock)
			got, err := u.Create(context.Background(), test.input)
			assert.Equal(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}