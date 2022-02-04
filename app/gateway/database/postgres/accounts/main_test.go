package account

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest"

	"stoneBanking/app/gateway/database/postgres"
	"stoneBanking/app/gateway/database/postgres/pgtest"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	os.Exit(SetupTests(m))
}

func SetupTests(m *testing.M) int {
	dbName := pgtest.GetRandomDBName()
	pool, err := dockertest.NewPool("")
	ctx := context.Background()
	if err != nil {
		log.Fatalf("error when trying to connect to docker")
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=" + dbName,
			"listen_addresses = '*'",
		},
	})
	if err != nil {
		log.Fatalf("has not possible to connect to resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/%s?sslmode=disable", hostAndPort, dbName)
	resource.Expire(120) //nolint: errorlint
	pool.MaxWait = 120 * time.Second

	if err = pool.Retry(func() error {
		testPool, err := pgx.Connect(ctx, dbUrl)
		if err != nil {
			return err
		}
		return testPool.Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	newConn, err := pgxpool.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	testPool = newConn
	defer teardown()

	return m.Run()
}

func SetDatabase(t *testing.T, dbName string) *pgxpool.Pool {
	err := CreateDB(dbName, testPool)
	if err != nil {
		log.Fatalf("has not possible to create the new database error:%s", err.Error())
	}
	conn := testPool
	dbUrl := strings.Replace(conn.Config().ConnString(), conn.Config().ConnConfig.Database, dbName, 1)
	pgxConn, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("has not possible to connect to database %v", err.Error())
	}
	migrationPath := "file:../migrations"

	err = postgres.Migrate(migrationPath, dbUrl)
	if err != nil {
		log.Fatalf("error during migration %v", err)
	}
	return pgxConn
}

func GetConn() *pgxpool.Pool {
	return testPool
}

func TruncateTable(ctx context.Context, db *pgx.Conn) error {
	sqlQuery := `truncate accounts RESTART IDENTITY cascade`
	_, err := db.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}

func CreateDB(dbName string, conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), fmt.Sprintf(`CREATE DATABASE %s`, dbName))
	return err
}

func teardown() {
	testPool.Close()
}
