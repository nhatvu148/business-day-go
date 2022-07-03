package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/nhatvu148/business-day-go/db"
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

	conn := db.InitDB()
	defer conn.Close()

	r := http.NewServeMux()
	r.HandleFunc("/", HomePageHandler)
	r.HandleFunc("/catfact", CatFactPageHandler)
	r.HandleFunc("/business-day", BusinessDayHandler)
	r.HandleFunc("/custom-holiday", CustomHolidayHandler)

	m := middlewares.RequestPathLogger(r)
	log.Fatal().Err(http.ListenAndServe(os.Getenv("PORT"), m)).Msg("")
}

func CustomHolidayHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		dateString := query.Get("date")
		_, date := tools.IsValidDate(dateString)
		customHolidays, err := tools.DB.GetCustomHolidays(date)

		if err != nil {
			log.Err(err).Msg("Get Custom holiday error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(customHolidays)
		if err != nil {
			log.Err(err).Msg("JSON marshal error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(res)

	// case "POST":

	default:
		log.Err(&CustomError{msg: "Unsupported method"}).Msg("Unsupported method")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	distPath := fmt.Sprintf("%s/client/dist", tools.RootPath)
	htmlPath := fmt.Sprintf("%s/client/dist/index.html", tools.RootPath)

	fileServer := http.FileServer(http.Dir(distPath))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)

	if !fileMatcher.MatchString(r.URL.Path) {
		http.ServeFile(w, r, htmlPath)
	} else {
		fileServer.ServeHTTP(w, r)
	}
}

func CatFactPageHandler(w http.ResponseWriter, r *http.Request) {
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
