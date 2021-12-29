package transfer

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"stoneBanking/app/gateway/database/postgres"
	"testing"
	"time"

	"github.com/ory/dockertest"
)

var databaseTest *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Erro ao conectar ao docker")
	}
	resource := setupTests(*pool)

	defer dropTests(*pool, &resource)

	code := m.Run()
	os.Exit(code)
}

func setupTests(pool dockertest.Pool) dockertest.Resource {
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
		databaseTest, err := sql.Open("postgres", dbUrl)
		if err != nil {
			return err
		}
		return databaseTest.Ping()
	}); err != nil {
		log.Fatalf("Não foi possível conectar ao docker: %s", err)
	}
	setDatabase(*resource)
	return *resource
}

func setDatabase(resource dockertest.Resource) {
	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	databaseTest, _ = sql.Open("postgres", dbUrl)
	migrationPath := "file:../migrations"
	err := postgres.Migrate(migrationPath, dbUrl)
	if err != nil {
		log.Fatalf("erro na migração %v", err)
	}
}

func dropTests(pool dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Não foi possível limpar o banco - %s", err)
	}
}

func TruncateTable(db *sql.DB) error {
	sqlQuery := `TRUNCATE transfers`
	_, err := db.Exec(sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
