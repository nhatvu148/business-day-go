package main

import (
	"net/http"
	"os"
	"time"

	handlers "github.com/nhatvu148/business-day-go/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func handleRequests() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/business-day", handlers.IsBusinessDayHandler)
	log.Fatal().Err(http.ListenAndServe(":54528", nil)).Msg("")
}

func main() {
	logType := os.Getenv("LOG_TYPE")
	if logType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	handleRequests()
}
