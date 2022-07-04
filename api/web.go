package api

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

func (app *Application) Render(w http.ResponseWriter, t string) {
	partials := []string{
		fmt.Sprintf("%s/templates/base.html", app.Config.RootPath),
		fmt.Sprintf("%s/templates/header.html", app.Config.RootPath),
		fmt.Sprintf("%s/templates/footer.html", app.Config.RootPath),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/templates/%s", app.Config.RootPath, t))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		log.Err(err).Msg("Template file parsing error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	catFact, err := GetCatFact("https://catfact.ninja/fact")
	if err != nil {
		log.Err(err).Msg("GetCatFact error")
	}

	ip := Input{CurrentTime: time.Now().Format(time.RFC1123), CatFact: catFact}
	if err := tmpl.Execute(w, ip); err != nil {
		log.Err(err).Msg("Template execution error")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetCatFact(URL string) (string, error) {
	resp, err := http.Get(URL)
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
