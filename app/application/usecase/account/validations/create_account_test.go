package account

import (
	"database/sql"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"testing"

	"github.com/stretchr/testify/assert"
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
			name: "dados de conta corretos, deve passar pela validação",
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
			name: "dados de conta faltando nome, deve apresentar erro",
			input: account.Account{
				Name:    "",
				CPF:     "761.647.810-78",
				Secret:  "J0@0doR10",
				Balance: 0,
			},
			wantErr: customError.ErrorAccountNameRequired,
		},
		{
			name: "dados de conta faltando cpf, deve apresentar erro",
			input: account.Account{
				Name:    "Joao do Rio",
				CPF:     "",
				Secret:  "J0@0doR10",
				Balance: 0,
			},
			wantErr: customError.ErrorAccountCPFRequired,
		},
		{
			name: "dados de conta faltando a senha, deve apresentar erro",
			input: account.Account{
				Name:    "Joao do Rio",
				CPF:     "761.647.810-78",
				Secret:  "",
				Balance: 0,
			},
			wantErr: customError.ErrorAccountSecretRequired,
		},
		{
			name: "dados de conta apresentam saldo negativo, deve apresentar erro",
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
