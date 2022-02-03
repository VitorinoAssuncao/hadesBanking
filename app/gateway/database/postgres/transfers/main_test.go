package transfer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"stoneBanking/app/gateway/database/postgres"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/ory/dockertest"
)

var databaseTest *pgx.Conn

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("error when connecting to docker")
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
		log.Fatalf("has not possible to initialize the resource %s", err)
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
		log.Fatalf("has not possible to connect to docker: %s", err)
	}
	setDatabase(*resource)
	return *resource
}

func setDatabase(resource dockertest.Resource) {
	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	databaseTest, _ = pgx.Connect(context.Background(), dbUrl)
	migrationPath := "file:../migrations"
	err := postgres.Migrate(migrationPath, dbUrl)
	if err != nil {
		log.Fatalf("error during migration %v", err)
	}
}

func dropTests(pool dockertest.Pool, resource *dockertest.Resource) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("has not possible to drop the database - %s", err)
	}
}

func TruncateTable(ctx context.Context, db *pgx.Conn) error {
	sqlQuery := `TRUNCATE transfers, accounts RESTART IDENTITY cascade`
	_, err := db.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
