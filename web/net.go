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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {

	// GET Request
	data := WebData{
		Title:    "Hangman",
		Image:    hang.HangmanProgress,
		Progress: hang.CurrentWord,
	}

	// POST Request
	if r.Method == http.MethodPost {

		hang.Input = r.FormValue("letter")
		hang.WordInput = r.Form.Get("word")

		hang.Run()

		// DEBUG
		fmt.Println("current word: ", hang.WordProgress)
		fmt.Println("word to guess: ", hang.Word)
		fmt.Println("hangman pos: ", hang.HangmanPosition)
		fmt.Println("attempts: ", hang.Attempts)
	}

	hang.Template = "web/game.html"

	// HTML File parsing
	tmpl, err := template.ParseFiles(hang.Template)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		fmt.Println("Template path: ", hang.Template)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
