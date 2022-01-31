package account

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest"

	"stoneBanking/app/gateway/database/postgres"
)

var databaseTest *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("error when trying to connect to docker")
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
		log.Fatalf("has not possible to connect to resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	resource.Expire(120) //nolint: errorlint
	pool.MaxWait = 120 * time.Second

	if err = pool.Retry(func() error {
		databaseTest, err := sql.Open("postgres", dbUrl)
		if err != nil {
			return err
		}
		return databaseTest.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
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
		log.Fatalf("error during migration %v", err)
	}
}

func TruncateTable(db *sql.DB) error {
	sqlQuery := `truncate accounts RESTART IDENTITY cascade`
	_, err := db.Exec(sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
