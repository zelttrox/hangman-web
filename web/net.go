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
	http.HandleFunc("/reset", Resetgame)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.ListenAndServe(":8080", nil)
}

// Server handler
func Index(w http.ResponseWriter, r *http.Request) {

	// Use hang.Template directly
	data := WebData{
		Image:    hang.HangmanProgress,
		Progress: hang.CurrentWord,
		Attempts: hang.Attempts,
	}

	// POST Request
	if r.Method == http.MethodPost {

		hang.Input = r.FormValue("letter")
		hang.WordInput = r.Form.Get("word")

		hang.Run()

		data.Image = hang.HangmanProgress
		data.Progress = hang.CurrentWord
		data.Attempts = hang.Attempts

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

		
		return
	}

	LoadTemplate(w, data)
}

// Parses HTML template files
func LoadTemplate(w http.ResponseWriter, data WebData) {
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
func Resetgame(w http.ResponseWriter, r* http.Request) {
	
	hang.Reset()
	http.Redirect(w, r, "/", http.StatusSeeOther)

}