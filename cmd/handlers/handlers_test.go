package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"testing"

	handlers "github.com/nhatvu148/business-day-go/cmd/handlers"
	"github.com/rs/zerolog/log"
)

// Set the testing directory the same as project root
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Error().Err(err).Msg("Change directory error")
	}
}

func TestHomePageHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handlers.HomePageHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	htmlString := string(html)

	expectedString := "Welcome to Business Day API!"
	if !strings.Contains(htmlString, expectedString) {
		t.Errorf("html content does not contain the expected string: %v", expectedString)
	}

	re := regexp.MustCompile(`(Mon|Tue|Wed|Thu|Fri|Sat|Sun), \d\d [A-Za-z]{3} [\d]{4} \d\d:\d\d:\d\d UTC`)
	if !re.MatchString(htmlString) {
		t.Errorf("error displaying the current time")
	}
}

func TestIsBusinessDayHandler(t *testing.T) {
	checkBusinessDayResult := func(t testing.TB, date string, expected bool) {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/business-day?date=%s", date), nil)
		w := httptest.NewRecorder()

		handlers.BusinessDayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		var businessDayResult handlers.BusinessDayResult
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
		checkBusinessDayResult(t, "2023-01-01", false)
	})
}
