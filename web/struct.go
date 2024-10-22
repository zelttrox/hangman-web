package server

import (
	"net/http"
)

var Server http.Handler // Ficher du serveur
// var MethodType http.Request    // Type de request (GET ou POST)
// var Return http.ResponseWriter // Retour de la requête (GET ou POST)
var Template string // Fichier HTML utilisé pour afficher le site web
