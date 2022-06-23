package main

import (
	"net/http"
	"os"
	"time"

	handlers "github.com/nhatvu148/business-day-go/handlers"
	"github.com/nhatvu148/business-day-go/middlewares"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func handleRequests() {
	r := http.NewServeMux()

	r.HandleFunc("/", handlers.HomePageHandler)
	r.HandleFunc("/business-day", handlers.BusinessDayHandler)

	m := middlewares.RequestPathLogger(r)
	log.Fatal().Err(http.ListenAndServe(":54528", m)).Msg("")
}

func main() {
	logType := os.Getenv("LOG_TYPE")
	if logType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	handleRequests()
}
