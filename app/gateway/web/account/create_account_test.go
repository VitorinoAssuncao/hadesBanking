package account

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	usecase "stoneBanking/app/application/usecase/account"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/gateway/web/account/vo/input"
	"stoneBanking/app/gateway/web/account/vo/output"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	accountInput := input.CreateAccountVO{
		Name:    "Joao",
		CPF:     "761.647.810-78",
		Secret:  "J0@0doR10",
		Balance: 100,
	}

	body, _ := json.Marshal(accountInput)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))

	testCases := []struct {
		name            string
		accountUsecase  usecase.Usecase
		tokenRepository token.Repository
		input           account.Account
		wantCode        int
		wantResult      output.AccountOutputVO
		wantErr         error
	}{
		{
			name: "conta cadastrada com sucesso, quando dados corretos",
			accountUsecase: usecase.New(
				&account.RepositoryMock{
					CreateFunc: func(ctx context.Context, accountData account.Account) (account.Account, error) {
						return account.Account{
							ID:         1,
							Name:       "Joao do Rio",
							ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
							CPF:        "761.647.810-78",
							Secret:     "J0@0doR10",
							Balance:    0,
						}, nil
					},
					GetByCPFFunc: func(ctx context.Context, accountCPF string) (account.Account, error) {
						return account.Account{}, sql.ErrNoRows
					},
				},
				&token.RepositoryMock{}),
			tokenRepository: &token.RepositoryMock{},
			input: account.Account{
				ID:         1,
				Name:       "Joao do Rio",
				ExternalID: "94b9c27e-2880-42e3-8988-62dceb6b6463",
				CPF:        "761.647.810-78",
				Secret:     "J0@0doR10",
				Balance:    0,
			},
			wantCode: 200,
			wantResult: output.AccountOutputVO{
				ID:      "94b9c27e-2880-42e3-8988-62dceb6b6463",
				Name:    "Joao do Rio",
				CPF:     "761.647.810-78",
				Balance: 0,
			},

			wantErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			controller := New(test.accountUsecase, test.tokenRepository)
			controller.Create(rec, req)

			resultBody := output.AccountOutputVO{}
			err := json.Unmarshal(rec.Body.Bytes(), &resultBody)
			if err != nil {
				t.Errorf(err.Error())
			}

			test.wantResult.Created_At = resultBody.Created_At
			assert.Equal(t, test.wantResult, resultBody)
			assert.Equal(t, rec.Code, test.wantCode)
		})
	}

}
