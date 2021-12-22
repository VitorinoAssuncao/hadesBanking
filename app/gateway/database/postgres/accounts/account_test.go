package postgres_account

import (
	"context"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"testing"
	"time"
)

func Test_Create(t *testing.T) {
	t.Run("Conta criada com sucesso - Dados corretos", func(t *testing.T) {
		account := account.Account{
			ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
			Name:       "Joao da Silva",
			Cpf:        "38330499912",
			Balance:    10000,
			Created_at: time.Now(),
		}

		ctx := context.Background()
		database := databaseMock
		accountRepository := NewAccountRepository(database)
		result, err := accountRepository.Create(ctx, &account)

		if err != nil {
			t.Errorf("Problemas no cadastro %v", err)
		}

		if result != &account {
			t.Errorf("Problemas no cadastro")
		}
	})
	t.Run("Conta com erro - Dados Duplicados", func(t *testing.T) {
		account := account.Account{
			ID:         "d3280f8c-570a-450d-89f7-3509bc84980d",
			Name:       "Joao da Silva",
			Cpf:        "38330499912",
			Balance:    10000,
			Created_at: time.Now(),
		}

		ctx := context.Background()
		database := databaseMock
		accountRepository := NewAccountRepository(database)
		_, err := accountRepository.Create(ctx, &account)

		if err == nil {
			t.Errorf("Deveria ocorrer erro, pois dados estão duplicados")
		}
	})
}

func Test_GetAll(t *testing.T) {
	t.Run("Retorna todas as contas cadastradas", func(t *testing.T) {
		accountsResult := []account.Account{}
		ctx := context.Background()
		database := databaseMock
		accountRepository := NewAccountRepository(database)

		accountsResult, err := accountRepository.GetAll(ctx)

		if err != nil {
			t.Errorf("Erro ao buscar todas as contas: %v", err)
		}

		if accountsResult[0].ID != "d3280f8c-570a-450d-89f7-3509bc84980d" {
			t.Errorf("Conta diferente da cadastrada")
		}
	})
}

func Test_GetByCPF(t *testing.T) {
	t.Run("CPF Correto, conta localizada", func(t *testing.T) {
		cpfTarget := "38330499912"

		ctx := context.Background()
		database := databaseMock
		accountRepository := NewAccountRepository(database)

		accountResult, err := accountRepository.GetByCPF(ctx, cpfTarget)

		if err != nil {
			t.Errorf("Problemas no cadastro %v", err)
		}
		if accountResult.Cpf != cpfTarget {
			t.Errorf("conta não localizada")
		}
	})
	t.Run("CPF incorreto, conta não localizada", func(t *testing.T) {
		cpfTarget := "38330499999"
		ctx := context.Background()

		database := databaseMock
		accountRepository := NewAccountRepository(database)

		_, err := accountRepository.GetByCPF(ctx, cpfTarget)

		if err == nil {
			t.Errorf("CPF deveria não existir")
		}
	})
}

func Test_GetByID(t *testing.T) {
	t.Run("ID Correto, conta localizada", func(t *testing.T) {
		idTarget := types.AccountID("d3280f8c-570a-450d-89f7-3509bc84980d")
		ctx := context.Background()

		database := databaseMock
		accountRepository := NewAccountRepository(database)

		accountResult, err := accountRepository.GetByID(ctx, idTarget)

		if err != nil {
			t.Errorf("Problemas no cadastro %v", err)
		}
		if accountResult.ID != idTarget {
			t.Errorf("conta não localizada")
		}
	})
	t.Run("ID Incorreto, conta  não localizada", func(t *testing.T) {
		idTarget := types.AccountID("d3280f8c-570a-450d-89f7-3509bc849899")
		ctx := context.Background()

		database := databaseMock
		accountRepository := NewAccountRepository(database)

		_, err := accountRepository.GetByID(ctx, idTarget)

		if err == nil {
			t.Errorf("ID não existente, conta não deveria ser encontrada")
		}
	})
}
