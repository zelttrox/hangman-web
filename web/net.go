package server

import (
	"fmt"
	"net/http"
)

func CreateWebsite() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	fmt.Println("BRANCH TEST")
}
//i want a refund
