package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"

	"github.com/nhatvu148/business-day-go/middlewares"
	tools "github.com/nhatvu148/business-day-go/tools"
	"github.com/nhatvu148/business-day-go/web"
	"github.com/rs/zerolog/log"
)

type BusinessDayResult struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

type CustomError struct {
	msg string
}

func (m *CustomError) Error() string {
	return m.msg
}

func HandleRequests() {
	tools.SetLogger()

	r := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./client/dist"))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !fileMatcher.MatchString(r.URL.Path) {
			http.ServeFile(w, r, "./client/dist/index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	// r.HandleFunc("/", HomePageHandler)

	r.HandleFunc("/business-day", BusinessDayHandler)

	m := middlewares.RequestPathLogger(r)
	log.Fatal().Err(http.ListenAndServe(os.Getenv("PORT"), m)).Msg("")
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

		log.Err(&CustomError{msg: result.Error}).Msg(result.Error)
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

	w.Write(res)
}
