package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	server "github.com/nhatvu148/business-day-go"
	"github.com/rs/zerolog/log"
)

func TestIsBusinessDay(t *testing.T) {
	checkIsBusinessDayResult := func(t testing.TB, dateString string, expected bool) {
		date, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			log.Error().Err(err).Msg("")
		}
		got := server.IsBusinessDay(date)

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

func TestIsBusinessDayHandler(t *testing.T) {
	checkBusinessDayResult := func(t testing.TB, date string, expected bool) {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/business-day?date=%s", date), nil)
		w := httptest.NewRecorder()

		server.IsBusinessDayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		var businessDayResult server.BusinessDayResult
		err = json.Unmarshal(data, &businessDayResult)
		if err != nil {
			t.Errorf("%v", err)
		}

		if businessDayResult.Error == "Invalid date" {
			t.Error("Invalid date")
		} else if businessDayResult.Result != expected {
			t.Errorf("expected %v got %v", expected, businessDayResult.Result)
		}
	}

	t.Run("Test Case 1", func(t *testing.T) {
		checkBusinessDayResult(t, "2022-06-01", true)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		checkBusinessDayResult(t, "2022-06-05", false)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		checkBusinessDayResult(t, "2022-12-24", false)
	})

	t.Run("Test Case 4", func(t *testing.T) {
		checkBusinessDayResult(t, "abcde", false)
	})
}

func TestIsValidDate(t *testing.T) {
	checkValidDateResult := func(t testing.TB, date string, expected bool) {
		got, _ := server.IsValidDate(date)
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
