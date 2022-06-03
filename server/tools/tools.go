package tools

import (
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
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
