package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
)

type BusinessDayResult struct {
	Result bool `json:"result"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/business-day", getIsBusinessDay)
	log.Fatal(http.ListenAndServe(":54528", nil))
}

func getIsBusinessDay(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getIsBusinessDay")

	query := r.URL.Query()
	date := query.Get("date")

	result := BusinessDayResult{Result: IsBusinessDay(date)}
	res, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(200)
	w.Write([]byte(string(res)))
}

func IsBusinessDay(dateString string) bool {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		panic(err)
	}

	return !(date.Weekday() == time.Saturday || date.Weekday() == time.Sunday || holiday.IsHoliday(date))
}

func main() {
	handleRequests()
}
