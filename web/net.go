package server

import (
	"fmt"
	"net/http"
	"text/template"
)

// Create the web server
func CreateWebsite() {
	Template = "web/static/template.html"
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {
	data := WebData{
		Title: "Hangman",
		Image: "web/static/images/hangman6.png",
	}
	tmpl, err := template.ParseFiles(Template)
	if err != nil {
		fmt.Println("Error")
	}
	tmpl.Execute(w, data)
}
