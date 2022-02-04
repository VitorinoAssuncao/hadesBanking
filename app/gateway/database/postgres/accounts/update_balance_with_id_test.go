package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/domain/types"
)

func Test_UpdateBalance(t *testing.T) {
	ctx := context.Background()
	database := testPool
	accountRepository := NewAccountRepository(database)
	testCases := []struct {
		name       string
		want       bool
		runBefore  func() (value types.ExternalID)
		inputID    string
		inputValue int
		wantErr    error
	}{
		{
			name: "with a input data from a account that exist, update the balance and return without errors",
			runBefore: func() (value types.ExternalID) {
				input := account.Account{
					Name:    "Joao da Silva",
					CPF:     "38330499912",
					Balance: 10000,
				}
				created, err := accountRepository.Create(ctx, input)

				if err == nil {
					value = created.ExternalID
				}

				return value
			},
			want:       true,
			inputValue: 100,
			wantErr:    nil,
		},
		{
			name:       "when trying to update a account with a id that not exist, return a error",
			want:       false,
			inputID:    "d7aefc42-4467-434a-9690-da4367cd3a1d",
			inputValue: 100,
			wantErr:    customError.ErrorAccountIDNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if test.runBefore != nil {
				test.inputID = string(test.runBefore())
			}
			err := accountRepository.UpdateBalance(ctx, test.inputValue, types.ExternalID(test.inputID))
			assert.Equal(t, test.wantErr, err)
		})
	}
}
