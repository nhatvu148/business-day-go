package handlers

import (
	"encoding/json"
	"net/http"

	tools "github.com/nhatvu148/business-day-go/cmd/tools"
	"github.com/nhatvu148/business-day-go/cmd/web"
	"github.com/rs/zerolog/log"
)

type BusinessDayResult struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	web.Render(w, "test.html")
}

func BusinessDayHandler(w http.ResponseWriter, r *http.Request) {
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
			log.Error().Err(err).Msg("JSON Marshal Error")
		}

		w.WriteHeader(500)
		w.Write(jsonResp)
		return
	}

	result = BusinessDayResult{Result: tools.IsBusinessDay(date), Error: ""}
	res, err := json.Marshal(result)

	if err != nil {
		log.Error().Err(err).Msg("JSON Marshal Error")

		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(res)
}
