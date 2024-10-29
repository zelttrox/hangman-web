package server

import (
	"fmt"
	"net/http"
	"text/template"
)

func CreateWebsite() {
	Template = "web/template.html"
	http.HandleFunc("/", Index)       // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil) // On lance le serveur local sur le port 8080
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := WebData{
		Title: "Hangman League Of Legends",
		Image: "web/images/hangman6.png",
	}
	tmpl, err := template.ParseFiles(Template)
	if err != nil {
		fmt.Println("Error")
	}
	tmpl.Execute(w, data)
}
