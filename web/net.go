package server

import (
	"net/http"
)

func CreateWebsite() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
