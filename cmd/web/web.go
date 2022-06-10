package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Input struct {
	CurrentTime string
	CatFact     string
}

type CatFact struct {
	Fact   string
	Length int
}

func Render(w http.ResponseWriter, t string) {
	partials := []string{
		"./cmd/web/templates/base.html",
		"./cmd/web/templates/header.html",
		"./cmd/web/templates/footer.html",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		log.Error().Err(err).Msg("Template file parsing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	catFact, err := GetCatFact()
	if err != nil {
		log.Error().Err(err).Msg("GetCatFact error")
	}

	ip := Input{CurrentTime: time.Now().Format(time.RFC1123), CatFact: catFact}
	if err := tmpl.Execute(w, ip); err != nil {
		log.Error().Err(err).Msg("Template execution error")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetCatFact() (string, error) {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var fact CatFact
	json.Unmarshal(body, &fact)

	return fact.Fact, nil
}
