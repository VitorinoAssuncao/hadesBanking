package pgtest

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest"
)

var testPool *pgxpool.Pool

func SetupTests(m *testing.M) int {
	dbName := GetRandomDBName()
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
		log.Fatalf("it was not possible to connect to resource: %s", err)
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

func teardown() {
	testPool.Close()
}
