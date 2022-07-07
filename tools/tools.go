package tools

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	holiday "github.com/holiday-jp/holiday_jp-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "123456789"
	dbname   = "custom_holiday"
)

type Config struct {
	RootPath     string
	DatabaseURL  string
	DSN          string
	MigrationURL string
	Port         int
	LogType      string
	Env          string
}

func (config *Config) LoadConfig() {
	config.RootPath = os.Getenv("ROOT_PATH")
	config.DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	config.DatabaseURL = os.Getenv("DATABASE_URL")
	config.MigrationURL = os.Getenv("MIGRATION_URL")

	p, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Err(err).Msg("Convert string to int error")
	}
	config.Port = p
	config.LogType = os.Getenv("LOG_TYPE")
	config.Env = os.Getenv("ENV")
}

func IsValidDate(dateString string) (bool, time.Time) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return false, date
	}
	return true, date
}

func IsBusinessDay(date time.Time) bool {
	return !(date.Weekday() == time.Saturday || date.Weekday() == time.Sunday || holiday.IsHoliday(date))
}

func (config *Config) SetLogger() {
	if config.LogType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
}
