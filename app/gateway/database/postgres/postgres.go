package postgres

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var database *sql.DB

func InitiliazeDatabase() {
	const user = "postgres"
	const pass = "postgres"
	const base = "bank_stone"
	const dbUrl = "postgres://" + user + ":" + pass + "@0.0.0.0:5432/" + base + "?sslmode=disable"
	const migrationPath = "file://app/gateway/database/postgres/migrations"
	db, _ := sql.Open("postgres", dbUrl)
	err := Migrate(migrationPath, dbUrl)

	if err != nil {
		fmt.Println(err.Error())
	}
	database = db
}

func RetrieveConnection() *sql.DB {
	return database
}

func Migrate(migration_string, db_string string) error {
	migration, err := migrate.New(migration_string, db_string)

	if err != nil {
		return err
	}

	err = migration.Up()

	if err != nil {
		return err
	}

	return nil
}
