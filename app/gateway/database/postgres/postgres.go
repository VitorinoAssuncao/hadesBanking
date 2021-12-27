package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var database *sql.DB

func InitiliazeDatabase() {
	const migrationPath = "file://app/gateway/database/postgres/migrations"
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	base := os.Getenv("POSTGRES_BASE")

	dbUrl := "postgres://" + user + ":" + pass + "@0.0.0.0:5432/" + base + "?sslmode=disable"

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
