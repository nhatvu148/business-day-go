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
	Result bool `json:"result"`
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

func IsBusinessDayHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: getIsBusinessDay")

	query := r.URL.Query()
	date := query.Get("date")

	result := BusinessDayResult{Result: IsBusinessDay(date)}
	res, err := json.Marshal(result)

	if err != nil {
		log.Error().Err(err).Msg("")
	}

	w.WriteHeader(200)
	w.Write([]byte(string(res)))
}

func IsBusinessDay(dateString string) bool {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return !(date.Weekday() == time.Saturday || date.Weekday() == time.Sunday || holiday.IsHoliday(date))
}

func main() {
	logType := os.Getenv("LOG_TYPE")
	if logType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	handleRequests()
}
