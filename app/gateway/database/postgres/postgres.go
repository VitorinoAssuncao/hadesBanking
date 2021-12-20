package postgres

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
)

var database *pgx.Conn

func InitiliazeDatabase() {
	const user = "postgres"
	const pass = "postgres"
	const base = "bank_stone"
	const dbUrl = "postgres://" + user + ":" + pass + "@0.0.0.0:5432/" + base + "?sslmode=disable"
	connection, _ := pgx.Connect(context.Background(), dbUrl)
	err := Migrate(dbUrl)

	if err != nil {
		fmt.Println(err.Error())
	}
	database = connection
}

func RetrieveConnection() *pgx.Conn {
	return database
}

func Migrate(db_string string) error {
	path := "file://app/gateway/database/postgres/migrations"
	migration, err := migrate.New(path, db_string)

	if err != nil {
		return err
	}

	err = migration.Up()

	if err != nil {
		return err
	}

	return nil
}
