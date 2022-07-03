package db

import (
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/nhatvu148/business-day-go/models"
	"github.com/nhatvu148/business-day-go/tools"

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

func InitDB() *sql.DB {
	conn, err := OpenDB(tools.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection error")
	}
	tools.DB = models.DBModel{DB: conn}

	return conn
}
