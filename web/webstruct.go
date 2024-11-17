package server

type WebData struct {
	Title            string   // Page title
	Image            string   // State image of Hangman
	WordChosen       string   // Word to guess
	AttemptedLetters []string // Letters attempted array
}
