package server

import (
	"fmt"
	"net/http"
	"text/template"
)

// Create the web server
func CreateWebsite() {
	Template = "web/template.html"
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {
	data := WebData{
		Title: "Hangman",
		Image: "web/images/hangman6.png",
	}
	tmpl, err := template.ParseFiles(Template)
	if err != nil {
		fmt.Println("Error")
	}
	tmpl.Execute(w, data)
}
