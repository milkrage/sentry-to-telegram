package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
)

func Migrate(conn *sql.DB, path string) error {
	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}

	instance, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", path), "sqlite", driver)
	if err != nil {
		return fmt.Errorf("failed to create driver instance: %w", err)
	}

	err = instance.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
