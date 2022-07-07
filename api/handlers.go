package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/nhatvu148/business-day-go/models"
	"github.com/nhatvu148/business-day-go/tools"
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

type CustomHoliday struct {
	ID       int64  `json:"id"`
	Date     string `json:"date"`
	Category string `json:"category"`
}

type IDResponse struct {
	ID int64 `json:"id"`
}

func (app *Application) CustomHolidayHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		dateString := query.Get("date")
		if dateString == "" {
			// Get all Custom holidays
			customHolidays, err := app.DB.GetCustomHolidays()

			if err != nil {
				log.Err(err).Msg("Get Custom holidays error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Prevent Json Marshal from returning null when customHolidays is []
			customHolidays1 := make([]models.CustomHoliday, 0)
			if len(customHolidays) != 0 {
				customHolidays1 = customHolidays
			}

			res, err := json.Marshal(customHolidays1)
			if err != nil {
				log.Err(err).Msg("JSON marshal error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		} else {
			isDateValid, date := tools.IsValidDate(dateString)

			if !isDateValid {
				log.Err(&CustomError{msg: "Invalid date"}).Msg("Invalid date")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Get Custom holiday by date
			customHoliday, err := app.DB.GetCustomHolidayByDate(date)

			if err != nil {
				// How to get pq error code?
				var ErrNoRowsFound = errors.New("sql: no rows in result set")
				if err == ErrNoRowsFound {
					log.Err(err).Msg("Date not found")
					w.WriteHeader(http.StatusNotFound)
				} else {
					log.Err(err).Msg("Get Custom holiday by date error")
					w.WriteHeader(http.StatusInternalServerError)
				}
				return
			}

			res, err := json.Marshal(customHoliday)
			if err != nil {
				log.Err(err).Msg("JSON marshal error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		}

	case http.MethodPost:
		payload := CustomHoliday{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Err(err).Msg("Read body error")
		}

		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Err(err).Msg("Unmarshal error")
		}

		dateString := payload.Date
		category := payload.Category
		isDateValid, date := tools.IsValidDate(dateString)
		if !isDateValid {
			log.Err(&CustomError{msg: "Invalid date"}).Msg("Invalid date")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		customHoliday := models.CustomHoliday{Date: date, Category: category}

		id, err := app.DB.AddCustomHoliday(customHoliday)
		if err != nil {
			log.Err(err).Msg("Insert customHoliday error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResp, err := json.Marshal(IDResponse{ID: id})
		if err != nil {
			log.Err(err).Msg("JSON marshal error")
		}

		// WriteHeader before Write body to avoid superfluous response.WriteHeader call
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)

	case http.MethodPut:
		payload := CustomHoliday{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Err(err).Msg("Read body error")
		}

		err = json.Unmarshal(body, &payload)
		if err != nil {
			log.Err(err).Msg("Unmarshal error")
		}

		id := payload.ID
		dateString := payload.Date
		category := payload.Category
		isDateValid, date := tools.IsValidDate(dateString)
		if !isDateValid {
			log.Err(&CustomError{msg: "Invalid date"}).Msg("Invalid date")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		customHoliday := models.CustomHoliday{ID: id, Date: date, Category: category}

		err = app.DB.UpdateCustomHolidayById(customHoliday)
		if err != nil {
			// How to get pq error code?
			if err.Error() == "sql: no rows in result set" {
				log.Err(err).Msg("Date not found")
				w.WriteHeader(http.StatusNotFound)
			} else {
				log.Err(err).Msg("Update customHoliday error")
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

	case http.MethodDelete:
		query := r.URL.Query()
		dateString := query.Get("date")

		if dateString == "" {
			err := app.DB.DeleteAllCustomHoliday()

			if err != nil {
				log.Err(err).Msg("Delete All Custom holiday error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			isDateValid, date := tools.IsValidDate(dateString)

			if !isDateValid {
				log.Err(&CustomError{msg: "Invalid date"}).Msg("Invalid date")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err := app.DB.DeleteCustomHolidayBDate(date)

			if err != nil {
				log.Err(err).Msg("Delete Custom holiday by date error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	default:
		log.Err(&CustomError{msg: "Unsupported method"}).Msg("Unsupported method")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *Application) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	distPath := fmt.Sprintf("%s/client/dist", app.Config.RootPath)
	htmlPath := fmt.Sprintf("%s/client/dist/index.html", app.Config.RootPath)

	fileServer := http.FileServer(http.Dir(distPath))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)

	if !fileMatcher.MatchString(r.URL.Path) {
		http.ServeFile(w, r, htmlPath)
	} else {
		fileServer.ServeHTTP(w, r)
	}
}

func (app *Application) CatFactPageHandler(w http.ResponseWriter, r *http.Request) {
	app.Render(w, "content.html")
}

func (app *Application) BusinessDayHandler(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		return
	}

	isBusinessDay := false
	customHoliday, err := app.DB.GetCustomHolidayByDate(date)
	if err != nil {
		isBusinessDay = tools.IsBusinessDay(date)
	} else if customHoliday.Category == "Business day" {
		isBusinessDay = true
	} else if customHoliday.Category == "Holiday" {
		isBusinessDay = false
	}

	result = BusinessDayResult{Result: isBusinessDay, Error: ""}
	res, err := json.Marshal(result)

	if err != nil {
		log.Err(err).Msg("JSON marshal error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
