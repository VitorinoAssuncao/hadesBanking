package pgtest

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"stoneBanking/app/gateway/database/postgres"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetRandomDBName() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return fmt.Sprintf("db_%d", n)
}

func SetDatabase(dbName string) (*pgxpool.Pool, error) {
	err := createDB(dbName, testPool)
	if err != nil {
		return nil, fmt.Errorf("was not possible to create the new database error:%s", err.Error())
	}
	conn := testPool
	dbUrl := strings.Replace(conn.Config().ConnString(), conn.Config().ConnConfig.Database, dbName, 1)
	pgxConn, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("was not possible to connect to database %v", err.Error())
	}
	migrationPath := "file:../migrations"

	err = postgres.Migrate(migrationPath, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error during migration %v", err)
	}

	return pgxConn, nil
}

func createDB(dbName string, conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), fmt.Sprintf(`CREATE DATABASE %s`, dbName))
	return err
}

func dropDB(dbName string, conn *pgxpool.Pool) {
	_, err := conn.Exec(context.Background(), fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, dbName))
	if err != nil {
		log.Fatalf("was not possible to delete the database %s", err.Error())
	}
}
