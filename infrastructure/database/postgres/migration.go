package postgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var migrationFilesPath = "file://./migrations"

func executeMigrations(conn *sql.DB) {
	dbname := os.Getenv("DATABASE_NAME")

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("error while starting migration driver config for db %s: %s\n", dbname, err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationFilesPath, dbname, driver)
	if err != nil {
		log.Fatalf("error while instantiate migration driver for db %s: %s\n", dbname, err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("error while executing migration for db %s: %s\n", dbname, err)
	}
}
