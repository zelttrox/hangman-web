package server

var Template string

type WebData struct {
	Title            string   // Page title
	Image            string   // State image of Hangman
	AttemptedLetters []string // Letters attempted array
}