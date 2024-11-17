package server

import (
	"fmt"
	hangman "main/scripts"
	"net/http"
	"text/template"
)

// Create the web server
func CreateWebsite() {
	Template = "web/win.html"
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {

	// GET Request
	data := WebData{
		Title:            "Hangman",
		Image:            hangman.HangmanProgress,
		WordChosen:       hangman.Word,
		AttemptedLetters: hangman.Attempted,
	}

	// HTML File parsing
	tmpl, err := template.ParseFiles(hangman.Template)
	if err != nil {
		fmt.Println("Error")
	}

	tmpl.Execute(w, data)

	// POST Request
	if r.Method == http.MethodPost {

		input := r.FormValue("letter")
		wordInput := r.Form.Get("word")

		hangman.Input = input
		hangman.WordInput = wordInput
	}
}
