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
			log.Err(err).Msg("")
		}
		got := tools.IsBusinessDay(date)

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	}

	cases := []struct {
		description string
		date        string
		expected    bool
	}{
		{
			description: "Test Case 1",
			date:        "2022-06-01",
			expected:    true,
		},
		{
			description: "Test Case 2",
			date:        "2022-06-05",
			expected:    false,
		},
		{
			description: "Test Case 3",
			date:        "2022-12-25",
			expected:    false,
		},
		{
			description: "Test Case 4",
			date:        "2023-01-01",
			expected:    false,
		},
		{
			description: "Test Case 5",
			date:        "2023-01-02",
			expected:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			checkIsBusinessDayResult(t, tt.date, tt.expected)
		})
	}
}

func TestIsValidDate(t *testing.T) {
	checkValidDateResult := func(t testing.TB, date string, expected bool) {
		got, _ := tools.IsValidDate(date)
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	}

	cases := []struct {
		description string
		date        string
		expected    bool
	}{
		{
			description: "Test Case 1",
			date:        "2022-06-01",
			expected:    true,
		},
		{
			description: "Test Case 2",
			date:        "2022-06-40",
			expected:    false,
		},
		{
			description: "Test Case 3",
			date:        "123456abc",
			expected:    false,
		},
		{
			description: "Test Case 4",
			date:        "2023-01-01",
			expected:    true,
		},
		{
			description: "Test Case 5",
			date:        "2023-01-02",
			expected:    true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			checkValidDateResult(t, tt.date, tt.expected)
		})
	}
}
