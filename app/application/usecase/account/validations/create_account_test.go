package account

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
)

func Test_ValidateAccountData(t *testing.T) {
	testCases := []struct {
		name      string
		input     account.Account
		runBefore func(db *sql.DB)
		want      account.Account
		wantErr   error
	}{
		{
			name: "with the right data, return nil",
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			wantErr: nil,
		},
		{
			name: "with the data missing the value of field 'Name' return a error",
			input: account.Account{
				Name:    "",
				CPF:     "761.647.810-78",
				Secret:  "J0@0doR10",
				Balance: 0,
			},
			wantErr: customError.ErrorAccountNameRequired,
		},
		{
			name: "with the data missing the value of field 'CPF' return a error",
			input: account.Account{
				Name:    "Joao do Rio",
				CPF:     "",
				Secret:  "J0@0doR10",
				Balance: 0,
			},
			wantErr: customError.ErrorAccountCPFRequired,
		},
		{
			name: "with the data missing the value of field 'Name' return a error",
			input: account.Account{
				Name:    "Joao do Rio",
				CPF:     "761.647.810-78",
				Secret:  "",
				Balance: 0,
			},
			wantErr: customError.ErrorAccountSecretRequired,
		},
		{
			name: "with the field 'Balance' having a value less than 0(zero), return a error",
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    -5,
			},
			wantErr: customError.ErrorAccountBalanceInvalid,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateAccountData(test.input)
			assert.Equal(t, err, test.wantErr)
		})
	}

}
