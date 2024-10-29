package server

import (
	"fmt"
	hg "main/scripts"
	"net/http"
	"text/template"
)

// Create the web server
func CreateWebsite() {
	Template = "web/template.html"
	http.HandleFunc("/", Index)
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
