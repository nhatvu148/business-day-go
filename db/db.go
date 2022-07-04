package db

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(dsn string) *sql.DB {
	conn, err := OpenDB(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection error")
	}

	return conn
}

func RunDBMigration(migrationURL string, databaseURL string) {
	migration, err := migrate.New(migrationURL, databaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}
}
