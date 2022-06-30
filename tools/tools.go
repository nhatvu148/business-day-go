package tools

import (
	"os"
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var RootPath = os.Getenv("ROOT_PATH")

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
