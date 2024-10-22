package server

import (
	"fmt"
	"net/http"
	"text/template"
)

func CreatetServer() {
	Template = "template.html"
	Server = http.FileServer(http.Dir("./web"))
	http.HandleFunc("/", RequestHandler)
}

func RequestHandler(writeTo http.ResponseWriter, requestType *http.Request) {

}

func RenderWebsite(writeTo http.ResponseWriter) {
	templateFile, err := template.ParseFiles(Template)
	if err != nil {
		fmt.Println("Could not load HTML template")
	}
	templateFile.Execute(writeTo)
	{

	}
}
