package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"stoneBanking/app/common/utils/config"
)

func InitializeDatabase(config config.Config) (*sql.DB, error) {
	const migrationPath = "file://app/gateway/database/postgres/migrations"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBBase, config.DBSSLMode)

	db, _ := sql.Open("postgres", dbUrl)
	err := Migrate(migrationPath, dbUrl)
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	return db, nil
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
