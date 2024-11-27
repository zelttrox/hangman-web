package server

import (
	"encoding/json"
	"fmt"
	hang "main/scripts"
	"net/http"
	"text/template"
)

// Create the web server
func CreateWebsite() {
	Template = hang.Template
	http.HandleFunc("/", Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {
	Template = hang.Template
	// GET Request
	data := WebData{
		Image:    hang.HangmanProgress,
		Progress: hang.CurrentWord,
	}

	// POST Request
	if r.Method == http.MethodPost {

		hang.Input = r.FormValue("letter")
		hang.WordInput = r.Form.Get("word")

		hang.Run()

		// DEBUG
		fmt.Println("attempts: ", hang.Attempts)
		fmt.Println("progress: ", hang.HangmanProgress)
		fmt.Println(Template)

		data.Image = hang.HangmanProgress
		data.Progress = hang.CurrentWord

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}

	hang.Template = "web/game.html"

	// HTML File parsing
	tmpl, err := template.ParseFiles(Template)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		fmt.Println("Template path: ", Template)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
