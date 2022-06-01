package main_test

import (
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
