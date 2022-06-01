package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/nhatvu148/business-day-go"
)

func TestIsBusinessDay(t *testing.T) {
	got := server.IsBusinessDay("2022-06-01")
	expected := true

	if got != expected {
		t.Errorf("expected '%v' but got '%v'", expected, got)
	}
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

		if businessDayResult.Result != expected {
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
}
