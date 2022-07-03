package tools

import (
	"fmt"
	"os"
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
	"github.com/nhatvu148/business-day-go/models"
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

var (
	DSN      = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB       = models.DBModel{DB: nil}
	RootPath = os.Getenv("ROOT_PATH")
)

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

func SetLogger() {
	logType := os.Getenv("LOG_TYPE")
	if logType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
}
