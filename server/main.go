package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BusinessDayResult struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Info().Msg("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/business-day", IsBusinessDayHandler)
	log.Fatal().Err(http.ListenAndServe(":54528", nil)).Msg("")
}

func IsValidDate(dateString string) (bool, time.Time) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return false, date
	}
	return true, date
}

func IsBusinessDayHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: getIsBusinessDay")

	query := r.URL.Query()
	dateString := query.Get("date")

	var result BusinessDayResult
	isDateValid, date := IsValidDate(dateString)
	if !isDateValid {
		result = BusinessDayResult{
			Result: false,
			Error:  "Invalid date",
		}
		jsonResp, err := json.Marshal(result)

		if err != nil {
			log.Error().Err(err).Msg("")
		}

		w.WriteHeader(500)
		w.Write(jsonResp)
		return
	}

	result = BusinessDayResult{Result: IsBusinessDay(date), Error: ""}
	res, err := json.Marshal(result)

	if err != nil {
		log.Error().Err(err).Msg("")
	}

	w.WriteHeader(200)
	w.Write([]byte(string(res)))
}

func IsBusinessDay(date time.Time) bool {
	return !(date.Weekday() == time.Saturday || date.Weekday() == time.Sunday || holiday.IsHoliday(date))
}

func main() {
	logType := os.Getenv("LOG_TYPE")
	if logType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	handleRequests()
}
