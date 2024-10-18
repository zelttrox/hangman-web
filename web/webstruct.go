package server

import (
	"net/http"
)

var Server http.Handler     // Ficher du serveur
var MethodType http.Request // Type de request (GET ou POST)
