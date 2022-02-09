package pgtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest"
)

var testPool *pgxpool.Pool

func SetupTests() (func(), error) {
	dbName := GetRandomDBName()
	pool, err := dockertest.NewPool("")
	ctx := context.Background()
	if err != nil {
		return nil, fmt.Errorf("error when trying to connect to docker: %s", err)
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
		return nil, fmt.Errorf("it was not possible to connect to resource: %s", err)
	}

	pool.MaxWait = 120 * time.Second
	_ = resource.Expire(120)
	hostAndPort := resource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user_name:secret@%s/%s?sslmode=disable", hostAndPort, dbName)

	if err = pool.Retry(func() error {
		testPool, err := pgx.Connect(ctx, dbUrl)
		if err != nil {
			return err
		}
		return testPool.Ping(ctx)
	}); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %s", err)
	}
	newConn, err := pgxpool.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	testPool = newConn

	teardownFN := func() {
		testPool.Close()
		dropDB(dbName, testPool)

	}
	return teardownFN, nil
}
