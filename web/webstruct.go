package server

var Template string

var MainPage string
var WinPage string
var LosePage string

type WebData struct {
	Title            string   // Page title
	Image            string   // State image of Hangman
	AttemptedLetters []string // Letters attempted array
}