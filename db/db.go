package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/walles/env"
)

var Connection *sql.DB

func Connect() {
	connection, err := sql.Open("postgres", env.MustGet("DB_URL", env.String))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Connection = connection
	runMigrations()
}

func runMigrations() {
	driver, err := postgres.WithInstance(Connection, &postgres.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get the db driver: %v\n", err)
		os.Exit(1)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not fetch migrations: %v\n", err)
		os.Exit(1)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Fprintf(os.Stderr, "Error while running migrations: %v\n", err)
		os.Exit(1)
	}
}
