package server

import (
	"fmt"
	hang "main/scripts"
	"net/http"
	"text/template"
)

// Create the web server
func CreateWebsite() {
	hang.Template = "web/game.html"
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {

	// GET Request
	data := WebData{
		Title:            "Hangman",
		Image:            hang.HangmanProgress,
		WordChosen:       hang.Word,
		AttemptedLetters: hang.AttemptedDisplay,
	}

	// HTML File parsing
	tmpl, err := template.ParseFiles(hang.Template)
	if err != nil {
		fmt.Println("Error")
	}

	tmpl.Execute(w, data)

	// POST Request
	if r.Method == http.MethodPost {

		input := r.FormValue("letter")
		wordInput := r.Form.Get("word")

		hang.Input = input
		hang.WordInput = wordInput
	}
}
