package tools_test

import (
	"testing"
	"time"

	tools "github.com/nhatvu148/business-day-go/tools"
	"github.com/rs/zerolog/log"
)

func TestIsBusinessDay(t *testing.T) {
	checkIsBusinessDayResult := func(t testing.TB, dateString string, expected bool) {
		date, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			log.Error().Err(err).Msg("")
		}
		got := tools.IsBusinessDay(date)

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	}

	t.Run("Test Case 1", func(t *testing.T) {
		checkIsBusinessDayResult(t, "2022-06-01", true)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		checkIsBusinessDayResult(t, "2022-06-05", false)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		checkIsBusinessDayResult(t, "2022-12-25", false)
	})
}

func TestIsValidDate(t *testing.T) {
	checkValidDateResult := func(t testing.TB, date string, expected bool) {
		got, _ := tools.IsValidDate(date)
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	}

	t.Run("Test Case 1", func(t *testing.T) {
		checkValidDateResult(t, "2022-06-01", true)
	})

	t.Run("Test Case 1", func(t *testing.T) {
		checkValidDateResult(t, "2022-06-40", false)
	})

	t.Run("Test Case 1", func(t *testing.T) {
		checkValidDateResult(t, "123456abc", false)
	})
}
