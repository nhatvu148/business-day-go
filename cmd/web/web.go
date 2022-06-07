package web

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Input struct {
	CurrentTime string
}

func Render(w http.ResponseWriter, t string) {
	ip := Input{CurrentTime: time.Now().Format(time.RFC1123)}

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ip); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
