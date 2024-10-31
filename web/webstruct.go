package server

var Template string

type WebData struct {
	Title            string   // Page title
	Image            string   // State image of Hangman
	Word string              // Mot choisi au hasard *
	Blankspace []string      // Nombre de cases du Hangman *
	IsHardcoreMode bool 	// Défini si le jeu est en mode hardcore ou en normal
}