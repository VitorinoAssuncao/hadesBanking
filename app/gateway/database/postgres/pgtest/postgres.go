package pgtest

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"stoneBanking/app/domain/entities/account"
	"stoneBanking/app/domain/entities/transfer"
	"stoneBanking/app/domain/types"
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

func CreateAccount(conn *pgxpool.Pool, acc account.Account) (accID types.ExternalID, err error) {
	sqlQuery :=
		`
			INSERT INTO
				accounts (name, cpf, secret, balance)
			VALUES
				($1, $2, $3, $4)
			RETURNING
				external_id
		`
	result, err := conn.Exec(context.Background(), sqlQuery, acc.Name, acc.CPF, acc.Secret, acc.Balance)
	if err != nil {
		return "", err
	}

	return types.ExternalID(result), nil
}

func CreateTransfer(conn *pgxpool.Pool, tf transfer.Transfer) (tfID types.ExternalID, err error) {
	sqlQuery :=
		`
		INSERT INTO
			transfers (account_origin_id, account_destiny_id, amount)
		VALUES
			($1, $2, $3)
		RETURNING
			id
		`
	result, err := conn.Exec(context.Background(), sqlQuery, tf.AccountOriginID, tf.AccountDestinationID, tf.Amount)
	if err != nil {
		return "", err
	}

	return types.ExternalID(result), nil
}
