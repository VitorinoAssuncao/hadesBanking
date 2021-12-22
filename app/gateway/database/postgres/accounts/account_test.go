package postgres_account

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/types"
	"stoneBanking/app/gateway/database/postgres"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
)

var databaseMock *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Erro ao conectar ao docker")
	}
	SetupTests(*pool)

	code := m.Run()

	os.Exit(code)
}

func SetupTests(pool dockertest.Pool) dockertest.Resource {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	})
	if err != nil {
		log.Fatalf("Não foi possível inicializar o recurso %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	resource.Expire(120)
	pool.MaxWait = 120 * time.Second

	if err = pool.Retry(func() error {
		DatabaseMock, err := sql.Open("postgres", dbUrl)
		if err != nil {
			return err
		}
		return DatabaseMock.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	setDatabase(*resource)
	return *resource
}

func setDatabase(resource dockertest.Resource) {
	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	databaseMock, _ = sql.Open("postgres", dbUrl)
	migrationPath := "file:../migrations"
	err := postgres.Migrate(migrationPath, dbUrl)
	if err != nil {
		log.Fatalf("erro na migração %v", err)
	}
}

func DropTests(pool dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Não foi possível limpar o banco - %s", err)
	}
}

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
		accountResult := &account.Account{}

		ctx := context.Background()
		database := databaseMock
		accountRepository := NewAccountRepository(database)

		accountResult, err := accountRepository.GetByCPF(ctx, cpfTarget, accountResult)

		if err != nil {
			t.Errorf("Problemas no cadastro %v", err)
		}
		if accountResult.Cpf != cpfTarget {
			t.Errorf("conta não localizada")
		}
	})
	t.Run("CPF incorreto, conta não localizada", func(t *testing.T) {
		cpfTarget := "38330499999"
		accountResult := &account.Account{}
		ctx := context.Background()

		database := databaseMock
		accountRepository := NewAccountRepository(database)

		_, err := accountRepository.GetByCPF(ctx, cpfTarget, accountResult)

		if err == nil {
			t.Errorf("CPF deveria não existir")
		}
	})
}

func Test_GetByID(t *testing.T) {
	t.Run("ID Correto, conta localizada", func(t *testing.T) {
		idTarget := types.AccountID("d3280f8c-570a-450d-89f7-3509bc84980d")
		accountResult := &account.Account{}
		ctx := context.Background()

		database := databaseMock
		accountRepository := NewAccountRepository(database)

		accountResult, err := accountRepository.GetByID(ctx, idTarget, accountResult)

		if err != nil {
			t.Errorf("Problemas no cadastro %v", err)
		}
		if accountResult.ID != idTarget {
			t.Errorf("conta não localizada")
		}
	})
	t.Run("ID Incorreto, conta  não localizada", func(t *testing.T) {
		idTarget := types.AccountID("d3280f8c-570a-450d-89f7-3509bc849899")
		accountResult := &account.Account{}
		ctx := context.Background()

		database := databaseMock
		accountRepository := NewAccountRepository(database)

		_, err := accountRepository.GetByID(ctx, idTarget, accountResult)

		if err == nil {
			t.Errorf("ID não existente, conta não deveria ser encontrada")
		}
	})
}
