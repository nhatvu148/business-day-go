package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/nhatvu148/business-day-go"
)

func TestIsBusinessDay(t *testing.T) {
	got := server.IsBusinessDay("2022-06-01")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIsBusinessDayHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/business-day?date=2022-06-01", nil)
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

	if businessDayResult.Result != false {
		t.Errorf("expected false got %v", businessDayResult.Result)
	}
}
