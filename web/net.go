package server

import (
	"net/http"
	"text/template"
)

func CreateWebsite() {
	Template = "template.html"
	templateFile := template.Must(template.ParseFiles(Template))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := WebData {
			Title: "Hangman",
		}

	}
