package handlers

import (
	"encoding/json"
	"net/http"

	tools "github.com/nhatvu148/business-day-go/tools"
	"github.com/nhatvu148/business-day-go/web"
	"github.com/rs/zerolog/log"
)

type BusinessDayResult struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	web.Render(w, "content.html")
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
			log.Err(err).Msg("JSON marshal error")
		}

		log.Error().Msg(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResp)
		return
	}

	result = BusinessDayResult{Result: tools.IsBusinessDay(date), Error: ""}
	res, err := json.Marshal(result)

	if err != nil {
		log.Err(err).Msg("JSON marshal error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
