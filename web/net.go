package server

import "net/http"

func CreateWebsite() {
	Server = http.FileServer(http.Dir("./web"))
	http.Handle("/web/", http.StripPrefix("/web/", Server))
}

func WebHandler() {
}