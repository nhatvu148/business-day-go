package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/nhatvu148/business-day-go/api"
)

func cleanupDB(t testing.TB, app *api.Application) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/custom-holiday", nil)
	w := httptest.NewRecorder()
	app.CustomHolidayHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %v got %v", http.StatusOK, res.StatusCode)
	}
}

func addCustomHoliday(t testing.TB, app *api.Application, date, category string) {
	bodyReader := strings.NewReader(fmt.Sprintf(`{"date": "%v", "category": "%v"}`, date, category))
	req := httptest.NewRequest(http.MethodPost, "/api/v1/custom-holiday", bodyReader)
	w := httptest.NewRecorder()
	app.CustomHolidayHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %v got %v", http.StatusCreated, res.StatusCode)
	}
}

func TestCustomHolidayHandler(t *testing.T) {
	app := api.SetupApp()
	var returningID int64 = 0

	t.Run("GET Empty Custom Holiday", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/custom-holiday", nil)
		w := httptest.NewRecorder()
		app.CustomHolidayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code %v got %v", http.StatusOK, res.StatusCode)
		}

		dataString := string(data)
		expectedString := `[]`
		if !strings.Contains(dataString, expectedString) {
			t.Errorf("expected %v got %v", expectedString, dataString)
		}

	})

	t.Run("POST Insert Custom Holiday by date", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"date": "2022-12-22", "category": "Business day"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/custom-holiday", bodyReader)
		w := httptest.NewRecorder()
		app.CustomHolidayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		var result api.CustomHoliday
		err = json.Unmarshal(data, &result)
		if err != nil {
			t.Errorf("%v", err)
		}

		if result.Id == 0 {
			t.Errorf("expected Id > 0 got %v", result.Id)
		}
		returningID = result.Id

		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status code %v got %v", http.StatusCreated, res.StatusCode)
		}
	})

	t.Run("PUT Update Custom Holiday by id", func(t *testing.T) {
		bodyReader := strings.NewReader(fmt.Sprintf(`{"id": %d, "date": "2022-12-24", "category": "Holiday"}`, returningID))
		req := httptest.NewRequest(http.MethodPut, "/api/v1/custom-holiday", bodyReader)
		w := httptest.NewRecorder()
		app.CustomHolidayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code %v got %v", http.StatusOK, res.StatusCode)
		}
	})

	t.Run("GET Custom Holiday by date", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/custom-holiday?date=2022-12-24", nil)
		w := httptest.NewRecorder()
		app.CustomHolidayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code %v got %v", http.StatusOK, res.StatusCode)
		}

		var result api.CustomHoliday
		err = json.Unmarshal(data, &result)
		if err != nil {
			t.Errorf("%v", err)
		}

		expectedString := "Holiday"
		if result.Category != expectedString {
			t.Errorf("expected %v got %v", expectedString, result.Category)
		}
	})

	t.Run("DELETE Custom Holiday by date", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/custom-holiday?date=2022-12-24", nil)
		w := httptest.NewRecorder()
		app.CustomHolidayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code %v got %v", http.StatusOK, res.StatusCode)
		}
	})

	cleanupDB(t, app)
}

func TestHomePageHandler(t *testing.T) {
	app := api.SetupApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	app.HomePageHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	htmlString := string(html)

	expectedString := `Custom Holidays`
	if !strings.Contains(htmlString, expectedString) {
		t.Errorf("html content does not contain the expected string: %v", expectedString)
	}
}

func TestCatFactPageHandler(t *testing.T) {
	app := api.SetupApp()

	req := httptest.NewRequest(http.MethodGet, "/catfact", nil)
	w := httptest.NewRecorder()
	app.CatFactPageHandler(w, req)
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
	app := api.SetupApp()

	checkBusinessDayResult := func(t testing.TB, date string, expected bool) {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/business-day?date=%s", date), nil)
		w := httptest.NewRecorder()

		app.BusinessDayHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		var businessDayResult api.BusinessDayResult
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

	addCustomHoliday(t, app, "2022-12-24", "Business day")
	addCustomHoliday(t, app, "2022-07-05", "Holiday")

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
			date:        "2022-12-24",
			expected:    true,
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
		{
			description: "Test Case 5",
			date:        "2022-07-05",
			expected:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			checkBusinessDayResult(t, tt.date, tt.expected)
		})
	}

	cleanupDB(t, app)
}
