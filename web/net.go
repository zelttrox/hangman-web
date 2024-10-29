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
		Title:            "Hangman",
		Image:            "https://img-cdn.pixlr.com/image-generator/history/65bb506dcb310754719cf81f/ede935de-1138-4f66-8ed7-44bd16efc709/medium.webp",
		AttemptedLetters: hg.AttemptedLetters,
	}
	tmpl, err := template.ParseFiles(Template)
	if err != nil {
		fmt.Println("Error")
	}
	tmpl.Execute(w, data)
}
