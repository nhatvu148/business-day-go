package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	tools "github.com/nhatvu148/business-day-go/tools"
	"github.com/rs/zerolog/log"
)

type BusinessDayResult struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Info().Msg("Endpoint Hit: homePage")
}

func IsBusinessDayHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Endpoint Hit: getIsBusinessDay")

	query := r.URL.Query()
	dateString := query.Get("date")

	var result BusinessDayResult
	isDateValid, date := tools.IsValidDate(dateString)
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

	result = BusinessDayResult{Result: tools.IsBusinessDay(date), Error: ""}
	res, err := json.Marshal(result)

	if err != nil {
		log.Error().Err(err).Msg("")
	}

	w.WriteHeader(200)
	w.Write([]byte(string(res)))
}
