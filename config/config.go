package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetConfig() {
	logType := os.Getenv("LOG_TYPE")
	if logType == "USER_FRIENDLY" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
}
