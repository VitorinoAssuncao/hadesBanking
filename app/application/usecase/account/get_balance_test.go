package account

import (
	"context"
	"stoneBanking/app/application/vo/output"
	"testing"
)

type RepositoryMock struct {
}

type UseCaseMock struct {
}

func (r RepositoryMock) GetBalance(ctx context.Context, accountID string) (output.AccountBalanceVO, error) {
	return output.AccountBalanceVO{Balance: 100}, nil
}

func Test_GetBalance(t *testing.T) {
	t.Run("Retorno de Saldo para conta existente", func(t *testing.T) {
	})
}
